package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"

	"github.com/rs/zerolog/log"
)

type process struct {
	Cmd    *exec.Cmd
	In     io.WriteCloser
	closer chan struct{}
}

func newProcess(out chan<- string, args ...string) (*process, error) {
	p := &process{}

	p.Cmd = exec.Command(args[0], args[1:]...) // nolint: gosec

	stdout, err := p.Cmd.StdoutPipe()
	if err != nil {
		log.Error().Err(err).Str("cmd", args[0]).Msg("failed to pipe out")
		return nil, err
	}

	stdin, err := p.Cmd.StdinPipe()
	if err != nil {
		log.Error().Err(err).Str("cmd", args[0]).Msg("failed to pipe in")
		return nil, err
	}
	p.In = stdin

	p.closer = make(chan struct{}, 1)
	go func() {
		defer stdout.Close()
		defer stdin.Close()
		r := bufio.NewReader(stdout)
		for {
			select {
			case <-p.closer:
				return
			default:
			}
			s, err := r.ReadString('\n')
			if err == io.EOF {
				continue
			}
			if err != nil {
				log.Error().Err(err).Msgf("failed to read out %s", args[0])
				return
			}
			out <- s
		}
	}()

	return p, p.Cmd.Start()
}

func (p *process) close() error {
	p.closer <- struct{}{}
	return p.Cmd.Process.Kill()
}

// LogAnalyzer receives log and analyze them with an Expect function.
type LogAnalyzer struct {
	c chan string

	Processes map[string]*process
}

// NewLogAnalyzer returns a new valid log analyzer.
func NewLogAnalyzer() *LogAnalyzer {
	return &LogAnalyzer{
		c:         make(chan string, 1000),
		Processes: make(map[string]*process, 0),
	}
}

// Close kill all pipe processes started with Cmd method.
func (a *LogAnalyzer) Close() {
	for key, p := range a.Processes {
		if err := p.close(); err != nil {
			log.Error().Err(err).Str("cmd", key).Msg("failed to kill process")
		}
	}
}

// NewProcess launch a new process with log plugged on log analyzer.
func (a *LogAnalyzer) NewProcess(name string, args ...string) error {
	p, err := newProcess(a.c, args...)
	if err != nil {
		return err
	}
	a.Processes[name] = p
	return nil
}

// Expect sends log into f and return error if f fail. Returns nil when f returns ok.
func (a *LogAnalyzer) Expect(f func(string) (bool, error)) error {
	for s := range a.c {
		fmt.Println(s)
		ok, err := f(s)
		if err != nil {
			return err
		}
		if ok {
			return nil
		}
	}
	return nil
}
