package models

type Validable interface {
	IsValid() bool
}
