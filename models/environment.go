package models

import (
	"github.com/ylascombe/go-api/gorm_custom"
)

type Environment struct {
	gorm_custom.GormModelCustom

	Name string `gorm:"not null;unique"`
}
