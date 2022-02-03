package server

import (
	"github.com/andy-a-o/rox-go/v5/core/entities"
	"github.com/andy-a-o/rox-go/v5/core/model"
)

type RoxString = model.RoxString

func NewRoxString(defaultValue string, options []string) RoxString {
	return entities.NewRoxString(defaultValue, options)
}
