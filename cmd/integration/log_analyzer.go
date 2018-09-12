package main

import (
	"bufio"
	"io"
	"os/exec"

	"github.com/rs/zerolog/log"
)

type bin struct {
	Cmd    *exec.Cmd
	In     io.WriteCloser
	closer chan struct{}
}

func (b *bin) run(out chan string, args ...string) error {

	b.Cmd = exec.Command(args[0], args[1:]...) // nolint: gas

	stdout, err := b.Cmd.StdoutPipe()
	if err != nil {
		log.Error().Err(err).Str("cmd", args[0]).Msg("failed to pipe out")
		return err
	}

	stdin, err := b.Cmd.StdinPipe()
	if err != nil {
		log.Error().Err(err).Str("cmd", args[0]).Msg("failed to pipe in")
		return err
	}
	b.In = stdin

	b.closer = make(chan struct{}, 0)
	go func() {
		defer stdout.Close()
		defer stdin.Close()
		r := bufio.NewReader(stdout)
		for {
			select {
			case <-b.closer:
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

	return b.Cmd.Start()
}

func (b *bin) close() error {
	if err := b.Cmd.Process.Kill(); err != nil {
		log.Error().Err(err).Str("cmd", b.Cmd.Path).Msg("failed to kill process")
	}

	return nil
}

// LogAnalyzer receives log and analyze them with an Expect function.
type LogAnalyzer struct {
	c chan string

	cmds []*exec.Cmd
}

// NewLogAnalyzer returns a new valid log analyzer.
func NewLogAnalyzer() *LogAnalyzer {
	return &LogAnalyzer{
		c: make(chan string, 1000),
	}
}

// Close kill all pipe processes started with Cmd method.
func (a *LogAnalyzer) Close() {
	for _, cmd := range a.cmds {
		if cmd == nil || cmd.Process == nil {
			continue
		}
		if err := cmd.Process.Kill(); err != nil {
			log.Error().Err(err).Str("cmd", cmd.Path).Msg("failed to kill process")
		}
	}
}

// Cmd runs a cmd and plug output (stdout) in analyzer chan.
func (a *LogAnalyzer) Cmd(args ...string) error {
	cmd := exec.Command(args[0], args[1:]...) // nolint: gas
	a.cmds = append(a.cmds, cmd)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Error().Err(err).Str("cmd", args[0]).Msg("failed to pipe out")
		return err
	}

	go func(stdout io.ReadCloser) {
		defer stdout.Close()
		r := bufio.NewReader(stdout)
		for {
			s, err := r.ReadString('\n')
			if err == io.EOF {
				continue
			}
			if err != nil {
				log.Error().Err(err).Msgf("failed to read out %s", args[0])
				return
			}
			a.c <- s
		}
	}(stdout)

	return cmd.Start()
}

// Expect sends log into f and return error if f fail. Returns nil when f returns ok.
func (a *LogAnalyzer) Expect(f func(string) (bool, error)) error {
	for s := range a.c {
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
