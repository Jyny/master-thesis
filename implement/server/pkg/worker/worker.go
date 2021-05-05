package worker

import (
	"io"
	"log"
	"os/exec"
	"server/pkg/model"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	maxWaiting = 100
	maxRunning = 2
)

type Task struct {
	MeetingID uuid.UUID
	Class     model.WorkerClass
	CMD       *exec.Cmd
}

type Worker struct {
	orm     *gorm.DB
	Waiting chan Task
	Running chan Task
}

func New() Worker {
	return Worker{
		Waiting: make(chan Task, maxWaiting),
		Running: make(chan Task, maxRunning),
	}
}

func (w *Worker) Start() {
	go w.listener()
}

func (w *Worker) listener() {
	for {
		task := <-w.Waiting

		worker := model.Worker{
			MeetingID: task.MeetingID,
			Class:     task.Class,
			Status:    model.PENDING,
		}
		err := w.orm.Create(&worker).Error
		if err != nil {
			log.Println(err)
		}

		w.Running <- task
		go w.runner(worker.ID)
	}
}

func (w *Worker) runner(workerID uuid.UUID) {
	task := <-w.Running

	stdout, err := task.CMD.StdoutPipe()
	if err != nil {
		log.Println(err)
	}
	stderr, err := task.CMD.StderrPipe()
	if err != nil {
		log.Println(err)
	}

	if err := task.CMD.Start(); err != nil {
		log.Println(err)
	}

	worker := model.Worker{
		Base: model.Base{
			ID: workerID,
		},
		Status: model.RUNNING,
	}
	err = w.orm.Model(&worker).Updates(&worker).Take(&worker).Error
	if err != nil {
		log.Println(err)
	}

	if err := task.CMD.Wait(); err != nil {
		log.Println(err)
	}

	stderrb, err := io.ReadAll(stderr)
	if err != nil {
		log.Println(err)
	}

	stdoutb, err := io.ReadAll(stdout)
	if err != nil {
		log.Println(err)
	}

	worker.StdOut = stdoutb
	worker.StdErr = stderrb
	worker.Status = model.COMPLETE
	err = w.orm.Model(&worker).Updates(&worker).Take(&worker).Error
	if err != nil {
		log.Println(err)
	}

	if worker.Class == model.ALIGN {
		shift, err := strconv.Atoi(string(stdoutb))
		if err != nil {
			log.Println(err)
		}

		w.Waiting <- Task{
			MeetingID: task.MeetingID,
			Class:     model.ANC,
			CMD:       exec.Command("python3", "anc.py", strconv.Itoa(shift), ""),
		}
	}
}
