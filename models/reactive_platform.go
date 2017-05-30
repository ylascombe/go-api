package models

type ReactivePlatform struct {
	Version   string `json:"version" yaml:"version"`
	ExtraVars map[string]string `json:"extra_vars" yaml:"extra_vars"`
	FeaturesStatus map[string]string `json:"features_status" yaml:"features_status"`
}