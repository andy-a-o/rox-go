package server

import (
	"github.com/andy-a-o/rox-go/v4/core/entities"
	"github.com/andy-a-o/rox-go/v4/core/model"
)

type RoxVariant = model.Variant

func NewRoxVariant(defaultValue string, options []string) RoxVariant {
	return entities.NewVariant(defaultValue, options)
}
