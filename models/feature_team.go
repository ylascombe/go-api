package models

import (
	"github.com/ylascombe/go-api/gorm_custom"
	"errors"
)

type FeatureTeam struct {
	gorm_custom.GormModelCustom

	Name      string `gorm:"not null;unique" json:"name" yaml:"name"`
	GitlabUrl string `gorm:"not null;unique" json:"gitlab_url" yaml:"gitlab_url"`
	GroupId   string `gorm:"not null;unique" json:"group_id" yaml:"group_id"`
}

type TransformedFeatureTeam struct {
	ID        uint       `gorm:"primary_key" yaml:"ID" json:"ID"`

	Name      string `gorm:"not null;unique" json:"name" yaml:"name"`
	GitlabUrl string `gorm:"not null;unique" json:"gitlab_url" yaml:"gitlab_url"`
	GroupId   string `gorm:"not null;unique" json:"group_id" yaml:"group_id"`
}

func (featureTeam FeatureTeam) IsValid() bool {
	return FeatureTeam{} != featureTeam && featureTeam.Name != ""
}

func NewFeatureTeam(name string) (*FeatureTeam, error) {
	featureTeam := FeatureTeam{Name: name}

	if featureTeam.IsValid() {
		return &featureTeam, nil
	} else {
		return nil, errors.New("Name is empty")
	}
}

func TransformFeatureTeam(featureTeam FeatureTeam) *TransformedFeatureTeam {
	return &TransformedFeatureTeam{
		ID: featureTeam.ID,
		Name: featureTeam.Name,
		GitlabUrl: featureTeam.GitlabUrl,
		GroupId: featureTeam.GroupId,
	}
}
