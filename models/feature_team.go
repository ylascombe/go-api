package models

import (
	"arc-api/gorm_custom"
	"errors"
)

type FeatureTeam struct {
	gorm_custom.GormModelCustom

	Name      string `gorm:"not null;unique" json:"name" yaml:"name"`
	GitlabUrl string `gorm:"not null;unique" json:"gitlab_url" yaml:"gitlab_url"`
	GroupId   string `gorm:"not null;unique" json:"group_id" yaml:"group_id"`
}

type TransformedFeatureTeam struct {
	ID        uint       `yaml:"ID" json:"ID"`

	Name      string `json:"name" yaml:"name"`
	GitlabUrl string `json:"gitlab_url" yaml:"gitlab_url"`
	GroupId   string `json:"group_id" yaml:"group_id"`
}

func (featureTeam FeatureTeam) IsValid() bool {
	return FeatureTeam{} != featureTeam &&
		featureTeam.Name != "" &&
		featureTeam.GroupId != "" &&
		featureTeam.GitlabUrl != ""
}

func NewFeatureTeam(name string, gitlabUrl string, groupId string) (*FeatureTeam, error) {
	featureTeam := FeatureTeam{Name: name, GitlabUrl: gitlabUrl, GroupId: groupId}

	if featureTeam.IsValid() {
		return &featureTeam, nil
	} else {
		return nil, errors.New("Name, GitlabUrl, or GroupId is empty")
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
