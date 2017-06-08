package models

import "github.com/jinzhu/gorm"

type AppModule struct {
	gorm.Model
	Artefact Artefact
	CommonConfig CommonConfig
}