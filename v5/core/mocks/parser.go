package mocks

import (
	"github.com/andy-a-o/rox-go/v5/core/context"
	"github.com/andy-a-o/rox-go/v5/core/roxx"
	"github.com/stretchr/testify/mock"
)

type Parser struct {
	mock.Mock
}

func (m *Parser) EvaluateExpression(expression string, context context.Context) roxx.EvaluationResult {
	args := m.Called(expression, context)
	return args.Get(0).(roxx.EvaluationResult)
}

func (m *Parser) AddOperator(name string, operation roxx.Operation) {
	m.Called(name, operation)
}
