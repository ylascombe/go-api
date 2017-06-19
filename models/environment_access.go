package models

import (
	"github.com/ylascombe/go-api/gorm_custom"
)

type EnvironmentAccess struct {
	gorm_custom.GormModelCustom

	ApiUser       ApiUser `gorm:"ForeignKey:ApiUserID"`
	ApiUserID     uint
	Environment   Environment `gorm:"ForeignKey:EnvironmentID"`
	EnvironmentID uint
}

type TransformedEnvironmentAccess struct {
	ID uint `json:"id"`

	TransformedApiUser       TransformedApiUser `gorm:"ForeignKey:ApiUserID"`
	ApiUserID     uint
	TransformedEnvironment   TransformedEnvironment `gorm:"ForeignKey:EnvironmentID"`
	EnvironmentID uint
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
