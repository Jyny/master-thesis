package model

import "github.com/google/uuid"

type Owner struct {
	Base
	PublicKey string
	SessionID uuid.UUID
	Session   Meeting
}
