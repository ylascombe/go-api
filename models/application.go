package models

type Application struct {
	Name  string `json:"name" yaml:"name"`
	AppModules []AppModule  `json:"appModules" yaml:"appModules"`
	FeatureTeam FeatureTeam `gorm:"ForeignKey:FeatureTeamID"`
	FeatureTeamID uint
}
