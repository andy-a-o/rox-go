package repositories_test

import (
	"github.com/andy-a-o/rox-go/v5/core/model"
	"github.com/andy-a-o/rox-go/v5/core/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExperimentRepositoryWillReturnNullWhenNotFound(t *testing.T) {
	exp := []*model.ExperimentModel{
		model.NewExperimentModel("1", "1", "1", false, []string{"a"}, nil),
	}

	expRepo := repositories.NewExperimentRepository()
	expRepo.SetExperiments(exp)

	assert.Nil(t, expRepo.GetExperimentByFlag("b"))
}

func TestExperimentRepositoryWillReturnWhenFound(t *testing.T) {
	exp := []*model.ExperimentModel{
		model.NewExperimentModel("1", "1", "1", false, []string{"a"}, nil),
	}

	expRepo := repositories.NewExperimentRepository()
	expRepo.SetExperiments(exp)

	assert.Equal(t, "1", expRepo.GetExperimentByFlag("a").ID)
}
