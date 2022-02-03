package server

import (
	"github.com/andy-a-o/rox-go/v5/core/entities"
	"github.com/andy-a-o/rox-go/v5/core/model"
)

type RoxFlag = model.Flag

func NewRoxFlag(defaultValue bool) RoxFlag {
	return entities.NewFlag(defaultValue)
}
