package models

import (
	"github.com/ylascombe/go-api/gorm_custom"
)

type Environment struct {
	gorm_custom.GormModelCustom

	Name string `gorm:"not null;unique"`
}


type TransformedEnvironment struct {
	ID uint `json:"id"`
	Name string `gorm:"not null;unique"`
}

func TransformEnvironment(environment Environment) *TransformedEnvironment {
	return &TransformedEnvironment{
		ID: environment.ID,
		Name: environment.Name,
	}
}
