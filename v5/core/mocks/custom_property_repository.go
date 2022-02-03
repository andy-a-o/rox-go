package mocks

import (
	"github.com/andy-a-o/rox-go/v5/core/model"
	"github.com/andy-a-o/rox-go/v5/core/properties"
	"github.com/stretchr/testify/mock"
)

type CustomPropertyRepository struct {
	mock.Mock
}

func (m *CustomPropertyRepository) AddCustomProperty(customProperty *properties.CustomProperty) {
	m.Called(customProperty)
}

func (m *CustomPropertyRepository) AddCustomPropertyIfNotExists(customProperty *properties.CustomProperty) {
	m.Called(customProperty)
}

func (m *CustomPropertyRepository) GetCustomProperty(name string) *properties.CustomProperty {
	args := m.Called(name)
	return args.Get(0).(*properties.CustomProperty)
}

func (m *CustomPropertyRepository) GetAllCustomProperties() []*properties.CustomProperty {
	args := m.Called()
	return args.Get(0).([]*properties.CustomProperty)
}

func (m *CustomPropertyRepository) RegisterPropertyAddedHandler(handler model.CustomPropertyAddedHandler) {
	m.Called(handler)
}
