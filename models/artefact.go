package models

import "github.com/jinzhu/gorm"

type Artefact struct {
	gorm.Model

	Name     string
	NexusUrl string
}
