package models

import (
	"arc-api/gorm_custom"
)

type AppModule struct {
	gorm_custom.GormModelCustom

	Artefact     Artefact
	CommonConfig CommonConfig
}
