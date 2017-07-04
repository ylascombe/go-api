package models

import (
	"errors"
	"time"
)

type Membership struct {
	// specify manually ID since it is a composed key
	UserID     uint `gorm:"primary_key"`
	FeatureTeamID uint `gorm:"primary_key"`
	CreatedAt     *time.Time `json:"-" yaml:"-"`
	UpdatedAt     *time.Time `json:"-" yaml:"-"`
	DeletedAt     *time.Time `sql:"index" json:"-" yaml:"-"`

	User       User `gorm:"ForeignKey:UserID"`
	FeatureTeam   FeatureTeam `gorm:"ForeignKey:FeatureTeamID"`
}

type TransformedMembership struct {
	UserID              uint `json:"user_id" yaml:"user_id"`
	FeatureTeamID          uint `json:"team_id" yaml:"team_id"`
	CreatedAt              *time.Time `json:"-" yaml:"-"`
	UpdatedAt              *time.Time `json:"-" yaml:"-"`
	DeletedAt              *time.Time `json:"-" yaml:"-"`

	TransformedUser     TransformedUser `json:"user" yaml:"user"`
	TransformedFeatureTeam TransformedFeatureTeam `json:"team" yaml:"team"`
}

func (membership Membership) IsValid() bool {
	return membership.UserID != 0 &&
		membership.User.ID == membership.UserID &&
		membership.FeatureTeamID != 0 &&
		membership.FeatureTeam.ID == membership.FeatureTeamID
}

func NewMembership(user User, featureTeam FeatureTeam) (*Membership, error) {
	if user.ID != 0 && featureTeam.ID != 0 {
		return &Membership{User: user, UserID: user.ID, FeatureTeam: featureTeam, FeatureTeamID: featureTeam.ID}, nil
	} else {
		return nil, errors.New("Invalid parameters when create Membership")
	}
}

func TransformMembership(membership Membership) *TransformedMembership {
	return &TransformedMembership{
		TransformedUser: *TransformUser(membership.User),
		UserID: membership.UserID,
		TransformedFeatureTeam: *TransformFeatureTeam(membership.FeatureTeam),
		FeatureTeamID: membership.FeatureTeamID,
	}
}
