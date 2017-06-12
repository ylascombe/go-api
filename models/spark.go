package models

type Spark struct {
	Version      string            `json:"version" yaml:"version"`
	ArtifactName string            `json:"artifact_name" yaml:"artifact_name"`
	ExtraVars    map[string]string `json:"extra_vars" yaml:"extra_vars"`
}
