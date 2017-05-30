package models

type Manifest struct {
	FormatVersion string           `json:"format_version" yaml:"format_version"`
	ReactPlatform ReactivePlatform `json:"reactive_platform" yaml:"reactive_platform"`
	Applications  []Application    `json:"applications" yaml:"applications"`
}
