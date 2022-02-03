package server

import (
	"github.com/andy-a-o/rox-go/v5/core/entities"
	"github.com/andy-a-o/rox-go/v5/core/model"
)

type RoxDouble = model.RoxDouble

func NewRoxDouble(defaultValue float64, options []float64) RoxDouble {
	return entities.NewRoxDouble(defaultValue, options)
}
