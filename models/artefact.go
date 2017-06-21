package models

import (
	"arc-api/gorm_custom"
)

type Artefact struct {
	gorm_custom.GormModelCustom

	Name     string
	NexusUrl string
	Version string
}
