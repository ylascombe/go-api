package models

import (
	"github.com/ylascombe/go-api/gorm_custom"
	"errors"
)

type ApiUser struct {
	gorm_custom.GormModelCustom

	Firstname    string `json:"firstname" yaml:"firstname"`
	Lastname     string `json:"lastname" yaml:"lastname"`
	Email        string `gorm:"not null;unique" json:"email" yaml:"email"`
	SshPublicKey string `json:"ssh_public_key" yaml:"ssh_public_key"`
	Pseudo       string `json:"pseudo" yaml:"pseudo"`
}

type TransformedApiUser struct {
	ID        uint       `gorm:"primary_key" yaml:"ID" json:"ID"`

	Firstname    string `json:"firstname" yaml:"firstname"`
	Lastname     string `json:"lastname" yaml:"lastname"`
	Email        string `gorm:"not null;unique" json:"email" yaml:"email"`
	SshPublicKey string `json:"ssh_public_key" yaml:"ssh_public_key"`
	Pseudo       string `json:"pseudo" yaml:"pseudo"`
}

func NewApiUser(firstname string, lastname string, pseudo string, email string, sshPublicKey string) (*ApiUser, error) {
	apiUser := ApiUser{Firstname: firstname, Lastname: lastname, Email: email, Pseudo: pseudo, SshPublicKey: sshPublicKey}

	if apiUser.IsValid() {
		return &apiUser, nil
	} else {
		return nil, errors.New("Given parameters are missing or invalid")
	}
}

func (apiUser ApiUser) IsValid() bool {
	return ApiUser{} != apiUser &&
		apiUser.Email != "" &&
		apiUser.Lastname != "" &&
		apiUser.Firstname != "" &&
		apiUser.Pseudo != ""
}

func TransformApiUser(apiUser ApiUser) *TransformedApiUser {
	return &TransformedApiUser{
		ID: apiUser.ID,
		Firstname: apiUser.Firstname,
		Lastname: apiUser.Lastname,
		Email: apiUser.Email,
		SshPublicKey: apiUser.SshPublicKey,
		Pseudo: apiUser.Pseudo,
	}
}
