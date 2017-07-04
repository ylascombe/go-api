package gorm_custom

import "time"

// XXX This Struct override gorm.Model one. Override it is required since the gorm.Model one does not specify how to JSON/YAML marshallize
// time.Time fields.
// In this version, fields are just ignored (thanks to the '-' key)

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models
//    type User struct {
//      gorm.Model
//    }
type GormModelCustom struct {
	ID        uint       `gorm:"primary_key" yaml:"ID" json:"ID"`
	CreatedAt *time.Time `json:"-" yaml:"-"`
	UpdatedAt *time.Time `json:"-" yaml:"-"`
	DeletedAt *time.Time `sql:"index" json:"-" yaml:"-"`
}
