package models

import (
	"arc-api/gorm_custom"
)

type EnvironmentAccess struct {
	gorm_custom.GormModelCustom

	ApiUser       ApiUser `gorm:"ForeignKey:ApiUserID"`
	ApiUserID     uint
	Environment   Environment `gorm:"ForeignKey:EnvironmentID"`
	EnvironmentID uint
}

type TransformedEnvironmentAccess struct {
	ID                     uint `json:"id"`

	TransformedApiUser     TransformedApiUser `json:"api_user" yaml:"api_user"`
	ApiUserID              uint        `json:"api_user_id" yaml:"api_user_id"`
	TransformedEnvironment TransformedEnvironment `json:"environment" yaml:"environment"`
	EnvironmentID          uint `json:"environment_id" yaml:"environment_id"`
}

func (envAccess EnvironmentAccess) IsValid() bool {
	return envAccess.ApiUserID != 0 &&
		envAccess.ApiUser.ID == envAccess.ApiUserID &&
		envAccess.EnvironmentID != 0 &&
		envAccess.Environment.ID == envAccess.EnvironmentID
}

func TransformEnvironmentAccess(envAccess EnvironmentAccess) *TransformedEnvironmentAccess {
	return &TransformedEnvironmentAccess{
		ID: envAccess.ID,
		TransformedApiUser: *TransformApiUser(envAccess.ApiUser),
		ApiUserID: envAccess.ApiUserID,
		TransformedEnvironment: *TransformEnvironment(envAccess.Environment),
		EnvironmentID: envAccess.EnvironmentID,
	}
}
