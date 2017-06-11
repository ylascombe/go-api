package models

import (
	"github.com/ylascombe/go-api/gorm_custom"
)

type ApiUser struct {
	gorm_custom.GormModelCustom

	Firstname string
	Lastname string
	Email string
	SshPublicKey string
}
