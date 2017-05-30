package models

type Application struct {
	Name  string `json:"name" yaml:"name"`
	Spark Spark  `json:"spark" yaml:"spark"`
	Api   Api    `json:"api" yaml:"api"`
}
