package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"estmanc/pkg/worker"
)

func main() {
	start := time.Now()
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

	durationEstm := time.Since(start)
	startANC := time.Now()

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

	durationANC := time.Since(startANC)
	durationTotal := time.Since(start)
	fmt.Println("time Estm : ", durationEstm)
	fmt.Println("time ANC  : ", durationANC)
	fmt.Println("time Total: ", durationTotal)
}
