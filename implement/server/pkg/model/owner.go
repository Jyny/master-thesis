package model

import "github.com/google/uuid"

type Owner struct {
	Base
	PublicKey string
	Challenge []byte
	Answer    []byte
	MeetingID uuid.UUID
	Meeting   Meeting
}
