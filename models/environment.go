package models

import (
	"arc-api/gorm_custom"
)

type Environment struct {
	gorm_custom.GormModelCustom

	Name string `gorm:"not null;unique" json:"name" yaml:"name"`
}

type TransformedEnvironment struct {
	ID   uint `json:"id"`
	Name string `json:"name" yaml:"name"`
}

func TransformEnvironment(environment Environment) *TransformedEnvironment {
	return &TransformedEnvironment{
		ID: environment.ID,
		Name: environment.Name,
	}
}
