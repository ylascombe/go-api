package models

import (
	"github.com/ylascombe/go-api/gorm_custom"
)

type Memberships struct {
	gorm_custom.GormModelCustom

	List       []Membership
}

func (memberships Memberships) IsValid() bool {
	return true
}

func NewMemberships() (*Memberships, error) {
	return &Memberships{}, nil
}