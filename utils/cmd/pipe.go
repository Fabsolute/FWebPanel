package cmd

import (
	"io"
	"os/exec"
)

type Pipe struct {
	pipes  []command
	stdOut io.Writer
}

type command struct {
	command   string
	arguments []string
}

func NewPipe(stdOut io.Writer) *Pipe {
	return &Pipe{make([]command, 0), stdOut}
}

func (p *Pipe) Pipe(cmd string, arguments ...string) *Pipe {
	p.pipes = append(p.pipes, command{cmd, arguments})
	return p
}

func (p *Pipe) Run() {
	if len(p.pipes) == 0 {
		return
	}

	commands := make([]*exec.Cmd, 0, len(p.pipes))

	command := exec.Command(p.pipes[0].command, p.pipes[0].arguments...)
	commands = append(commands, command)
	for index, pipe := range p.pipes {
		if index == 0 {
			continue
		}

		command2 := exec.Command(pipe.command, pipe.arguments...)
		command2.Stdin, _ = command.StdoutPipe()
		commands = append(commands, command2)
		command = command2
	}

	command.Stdout = p.stdOut
	for _, command := range commands {
		command.Start()
	}

	for _, command := range commands {
		command.Wait()
	}
}
