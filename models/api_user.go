package models

import (
	"github.com/ylascombe/go-api/gorm_custom"
)

type ApiUser struct {
	gorm_custom.GormModelCustom

	Firstname string `json:"firstname" yaml:"firstname"`
	Lastname string `json:"lastname" yaml:"lastname"`
	Email string `gorm:"not null;unique" json:"email" yaml:"email"`
	SshPublicKey string `json:"ssh_public_key" yaml:"ssh_public_key"`
}

func (apiUser ApiUser) IsValid() bool {
	return ApiUser{} != apiUser && apiUser.Email != "" && apiUser.Lastname != "" && apiUser.Firstname != ""
}
