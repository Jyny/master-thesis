package model

type Meeting struct {
	Base
	SessionKey    string
	EncSessionKey []byte
	AllowRegister bool
}
