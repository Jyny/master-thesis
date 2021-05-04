package model

import "github.com/google/uuid"

type Owner struct {
	Base
	PublicKey string
	MeetingID uuid.UUID
	Meeting   Meeting
}
