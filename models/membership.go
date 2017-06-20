package models

import (
	"errors"
	"time"
)

type Membership struct {
	// specify manually ID since it is a composed key
	ApiUserID     uint `gorm:"primary_key"`
	FeatureTeamID uint `gorm:"primary_key"`
	CreatedAt     *time.Time `json:"-" yaml:"-"`
	UpdatedAt     *time.Time `json:"-" yaml:"-"`
	DeletedAt     *time.Time `sql:"index" json:"-" yaml:"-"`

	ApiUser       ApiUser `gorm:"ForeignKey:ApiUserID"`
	FeatureTeam   FeatureTeam `gorm:"ForeignKey:FeatureTeamID"`
}

type TransformedMembership struct {
	ApiUserID     uint `json:"api_user_id" yaml:"api_user_id"`
	FeatureTeamID uint `json:"team_id" yaml:"team_id"`
	CreatedAt     *time.Time `json:"-" yaml:"-"`
	UpdatedAt     *time.Time `json:"-" yaml:"-"`
	DeletedAt     *time.Time `json:"-" yaml:"-"`

	TransformedApiUser       TransformedApiUser `json:"api_user" yaml:"api_user"`
	TransformedFeatureTeam   TransformedFeatureTeam `json:"team" yaml:"team"`
}

func (membership Membership) IsValid() bool {
	return membership.ApiUserID != 0 &&
		membership.ApiUser.ID == membership.ApiUserID &&
		membership.FeatureTeamID != 0 &&
		membership.FeatureTeam.ID == membership.FeatureTeamID
}

func NewMembership(apiUser ApiUser, featureTeam FeatureTeam) (*Membership, error) {
	if apiUser.ID != 0 && featureTeam.ID != 0 {
		return &Membership{ApiUser: apiUser, ApiUserID: apiUser.ID, FeatureTeam: featureTeam, FeatureTeamID: featureTeam.ID}, nil
	} else {
		return nil, errors.New("Invalid parameters when create Membership")
	}
}

func TransformMembership(membership Membership) *TransformedMembership {
	return &TransformedMembership{
		TransformedApiUser: *TransformApiUser(membership.ApiUser),
		ApiUserID: membership.ApiUserID,
		TransformedFeatureTeam: *TransformFeatureTeam(membership.FeatureTeam),
		FeatureTeamID: membership.FeatureTeamID,
	}
}
