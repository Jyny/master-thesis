package worker

import (
	"io"
	"log"
	"os/exec"
)

type Task struct {
	CMD *exec.Cmd
}

func (t *Task) Run() string {

	stdout, err := t.CMD.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err := t.CMD.StderrPipe()
	if err != nil {
		panic(err)
	}

	if err := t.CMD.Start(); err != nil {
		panic(err)
	}

	stderrb, err := io.ReadAll(stderr)
	if err != nil {
		panic(err)
	}

	stdoutb, err := io.ReadAll(stdout)
	if err != nil {
		panic(err)
	}

	if err := t.CMD.Wait(); err != nil {
		log.Println("stdout", string(stdoutb))
		log.Println("stderr", string(stderrb))
		panic(err)
	}

	log.Println("stdout", string(stdoutb))
	log.Println("stderr", string(stderrb))
	return string(stdoutb)
}
