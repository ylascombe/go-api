package models

import (
	"github.com/ylascombe/go-api/gorm_custom"
)

type Artefact struct {
	gorm_custom.GormModelCustom

	Name     string
	NexusUrl string
	Version string
}
