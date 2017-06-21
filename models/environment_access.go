package models

import (
	"arc-api/gorm_custom"
)

type EnvironmentAccess struct {
	gorm_custom.GormModelCustom

	User       User `gorm:"ForeignKey:UserID"`
	UserID     uint
	Environment   Environment `gorm:"ForeignKey:EnvironmentID"`
	EnvironmentID uint
}

type TransformedEnvironmentAccess struct {
	ID                     uint `json:"id"`

	TransformedUser     TransformedUser `json:"user" yaml:"user"`
	UserID              uint        `json:"user_id" yaml:"user_id"`
	TransformedEnvironment TransformedEnvironment `json:"environment" yaml:"environment"`
	EnvironmentID          uint `json:"environment_id" yaml:"environment_id"`
}

func (envAccess EnvironmentAccess) IsValid() bool {
	return envAccess.UserID != 0 &&
		envAccess.User.ID == envAccess.UserID &&
		envAccess.EnvironmentID != 0 &&
		envAccess.Environment.ID == envAccess.EnvironmentID
}

func TransformEnvironmentAccess(envAccess EnvironmentAccess) *TransformedEnvironmentAccess {
	return &TransformedEnvironmentAccess{
		ID: envAccess.ID,
		TransformedUser: *TransformUser(envAccess.User),
		UserID: envAccess.UserID,
		TransformedEnvironment: *TransformEnvironment(envAccess.Environment),
		EnvironmentID: envAccess.EnvironmentID,
	}
}
