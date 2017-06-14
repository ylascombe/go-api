package models

import (
	"github.com/ylascombe/go-api/gorm_custom"
)

type FeatureTeam struct {
	gorm_custom.GormModelCustom

	Name    string `json:"name" yaml:"name"`
}

func (featureTeam FeatureTeam) IsValid() bool {
	return FeatureTeam{} != featureTeam && featureTeam.Name != ""
}
