package models

import (
	"time"
)

type EnvironmentAccess struct {
	UserID        uint `gorm:"primary_key"`
	EnvironmentID uint `gorm:"primary_key"`
	CreatedAt     *time.Time `json:"-" yaml:"-"`
	UpdatedAt     *time.Time `json:"-" yaml:"-"`
	DeletedAt     *time.Time `sql:"index" json:"-" yaml:"-"`

	User          User `gorm:"ForeignKey:UserID"`
	Environment   Environment `gorm:"ForeignKey:EnvironmentID"`
}

type TransformedEnvironmentAccess struct {
	UserID                 uint `json:"user_id" yaml:"user_id"`
	EnvironmentID          uint `json:"environment_id" yaml:"environment_id"`

	TransformedUser        TransformedUser `json:"user" yaml:"user"`
	TransformedEnvironment TransformedEnvironment `json:"environment" yaml:"environment"`
}

func (envAccess EnvironmentAccess) IsValid() bool {
	return envAccess.UserID != 0 &&
		envAccess.User.ID == envAccess.UserID &&
		envAccess.EnvironmentID != 0 &&
		envAccess.Environment.ID == envAccess.EnvironmentID
}

func TransformEnvironmentAccess(envAccess EnvironmentAccess) *TransformedEnvironmentAccess {
	return &TransformedEnvironmentAccess{
		TransformedUser: *TransformUser(envAccess.User),
		UserID: envAccess.UserID,
		TransformedEnvironment: *TransformEnvironment(envAccess.Environment),
		EnvironmentID: envAccess.EnvironmentID,
	}
}
