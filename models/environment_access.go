package models

import (
	"github.com/ylascombe/go-api/gorm_custom"
)

type EnvironmentAccess struct {
	gorm_custom.GormModelCustom

	ApiUser ApiUser `gorm:"ForeignKey:ApiUserID"`
	ApiUserID uint
	Environment Environment `gorm:"ForeignKey:EnvironmentID"`
	EnvironmentID uint
}
