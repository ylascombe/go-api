package models

import (
	"github.com/ylascombe/go-api/gorm_custom"
)

type AppModule struct {
	gorm_custom.GormModelCustom

	Artefact     Artefact
	CommonConfig CommonConfig
}
