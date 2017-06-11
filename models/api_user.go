package models

import "github.com/jinzhu/gorm"

type ApiUser struct {
	gorm.Model
	Firstname string
	Lastname string
	Email string
	SshPublicKey string
}
