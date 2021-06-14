package main

import (
	"os"
	"os/exec"
	"strconv"
	"strings"

	"estmanc/pkg/worker"
)

func main() {
	if len(os.Args) < 3 {
		os.Exit(1)
	}

	estm := worker.Task{
		CMD: exec.Command("estimat-shift/main",
			os.Args[1],
			os.Args[2],
		),
	}
	shiftstr := estm.Run()
	shift, err := strconv.Atoi(strings.Trim(shiftstr, "\n"))
	if err != nil {
		panic(err)
	}

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dejam := "dejammed.wav"
	if len(os.Args) > 3 && os.Args[3] != "" {
		dejam = os.Args[3]
	}

	anc := worker.Task{
		CMD: exec.Command("docker",
			"run", "--rm", "-v", pwd+":/home/jovyan", "anc", "python3", "anc/anc.py",
			os.Args[1],
			os.Args[2],
			dejam,
			strconv.Itoa(shift),
		),
	}
	anc.Run()
}
