package model

import "github.com/google/uuid"

var (
	StatusType      = "status"
	WorkerClassType = "workerclass"
)

type Status string

const (
	PENDING  Status = "pending"
	RUNNING  Status = "running"
	COMPLETE Status = "complete"
)

type WorkerClass string

const (
	ALIGN WorkerClass = "align"
	ANC   WorkerClass = "anc"
)

type Worker struct {
	Base
	Class     WorkerClass
	Status    Status
	StdOut    []byte
	StdErr    []byte
	MeetingID uuid.UUID
	Meeting   Meeting
}
