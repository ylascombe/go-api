package models

import (
	"arc-api/gorm_custom"
	"errors"
)

type User struct {
	gorm_custom.GormModelCustom

	Firstname    string `json:"firstname" yaml:"firstname"`
	Lastname     string `json:"lastname" yaml:"lastname"`
	Email        string `gorm:"not null;unique" json:"email" yaml:"email"`
	SshPublicKey string `json:"ssh_public_key" yaml:"ssh_public_key"`
	Pseudo       string `json:"pseudo" yaml:"pseudo"`
}

type TransformedUser struct {
	ID        uint      `yaml:"ID" json:"ID"`

	Firstname    string `json:"firstname" yaml:"firstname"`
	Lastname     string `json:"lastname" yaml:"lastname"`
	Email        string `json:"email" yaml:"email"`
	SshPublicKey string `json:"ssh_public_key" yaml:"ssh_public_key"`
	Pseudo       string `json:"pseudo" yaml:"pseudo"`
}

func NewUser(firstname string, lastname string, pseudo string, email string, sshPublicKey string) (*User, error) {
	user := User{Firstname: firstname, Lastname: lastname, Email: email, Pseudo: pseudo, SshPublicKey: sshPublicKey}

	if user.IsValid() {
		return &user, nil
	} else {
		return nil, errors.New("Given parameters are missing or invalid")
	}
}

func (user User) IsValid() bool {
	return User{} != user &&
		user.Email != "" &&
		user.Lastname != "" &&
		user.Firstname != "" &&
		user.Pseudo != ""
}

func TransformUser(user User) *TransformedUser {
	return &TransformedUser{
		ID: user.ID,
		Firstname: user.Firstname,
		Lastname: user.Lastname,
		Email: user.Email,
		SshPublicKey: user.SshPublicKey,
		Pseudo: user.Pseudo,
	}
}
