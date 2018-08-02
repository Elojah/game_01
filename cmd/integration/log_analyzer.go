package main

import (
	"bufio"
	"os/exec"

	"github.com/rs/zerolog/log"
)

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
	cmd := exec.Command(args[0], args[1:]...)
	a.cmds = append(a.cmds, cmd)

	cmdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Error().Err(err).Str("cmd", args[0]).Msg("failed to pipe out")
		return err
	}

	go func() {
		r := bufio.NewReader(cmdout)
		for {
			s, err := r.ReadString('\n')
			if err != nil {
				log.Error().Err(err).Msgf("failed to read out %s", args[0])
				return
			}
			a.c <- s
		}
	}()

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
