package models

import "github.com/ylascombe/go-api/gorm_custom"

type CommonConfig struct {
	gorm_custom.GormModelCustom

	Json string
}
