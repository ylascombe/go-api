package models

import (
	"github.com/ylascombe/go-api/gorm_custom"
	"errors"
)

type Membership struct {
	gorm_custom.GormModelCustom

	ApiUser       ApiUser `gorm:"ForeignKey:ApiUserID"`
	ApiUserID     uint
	FeatureTeam   FeatureTeam `gorm:"ForeignKey:FeatureTeamID"`
	FeatureTeamID uint
}

func (membership Membership) IsValid() bool {
	return membership.ApiUserID != 0 &&
		membership.ApiUser.ID == membership.ApiUserID &&
		membership.FeatureTeamID != 0 &&
		membership.FeatureTeam.ID == membership.FeatureTeamID
}

func NewMembership(apiUser ApiUser, featureTeam FeatureTeam) (*Membership, error) {
	if apiUser.IsValid() && featureTeam.IsValid() {
		return &Membership{ApiUser: apiUser, FeatureTeam: featureTeam}, nil
	} else {
		return nil, errors.New("Invalid parameters")
	}
}