package models

import "github.com/ylascombe/go-api/gorm_custom"

type ReactivePlatform struct {
	gorm_custom.GormModelCustom

	Version        string            `json:"version" yaml:"version"`
	ExtraVars      string `json:"extra_vars" yaml:"extra_vars"`
	FeaturesStatus string `json:"features_status" yaml:"features_status"`
}
