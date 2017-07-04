package models

import "arc-api/gorm_custom"

type CommonConfig struct {
	gorm_custom.GormModelCustom

	Json string
}
