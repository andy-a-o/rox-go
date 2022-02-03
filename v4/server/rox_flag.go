package server

import (
	"github.com/andy-a-o/rox-go/v4/core/entities"
	"github.com/andy-a-o/rox-go/v4/core/model"
)

type RoxFlag = model.Flag

func NewRoxFlag(defaultValue bool) RoxFlag {
	return entities.NewFlag(defaultValue)
}
