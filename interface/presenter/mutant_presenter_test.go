package presenter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilmerpino/mutant/domain/model"
)

func TestResponseMutantsStats(t *testing.T) {
	presenter := &mutantPresenter{}

	expect := model.Stats{
		CountMutant: 2,
		CountHuman:  1,
		Ratio:       2,
	}
	count := []model.StatsResult{
		{
			Cant:     1,
			IsMutant: false,
		},
		{
			Cant:     2,
			IsMutant: true,
		},
	}

	actual := presenter.ResponseMutantsStats(count)
	assert.Equal(t, expect, actual)
}

func TestResponseMutantsStatsZero(t *testing.T) {
	presenter := &mutantPresenter{}

	expect := model.Stats{}
	var ac []model.StatsResult
	actual := presenter.ResponseMutantsStats(ac)
	assert.Equal(t, expect, actual)
}

func TestResponseMutantsZeroHuman(t *testing.T) {
	presenter := &mutantPresenter{}

	expect := model.Stats{
		CountMutant: 1,
		CountHuman:  0,
		Ratio:       0,
	}
	count := []model.StatsResult{
		{
			Cant:     1,
			IsMutant: true,
		},
	}

	actual := presenter.ResponseMutantsStats(count)
	assert.Equal(t, expect, actual)
}

func TestResponseMutantsZeroMutant(t *testing.T) {
	presenter := &mutantPresenter{}

	expect := model.Stats{
		CountMutant: 1,
		CountHuman:  1,
		Ratio:       1,
	}
	count := []model.StatsResult{
		{
			Cant:     1,
			IsMutant: true,
		},
		{
			Cant:     1,
			IsMutant: false,
		},
	}

	actual := presenter.ResponseMutantsStats(count)
	assert.Equal(t, expect, actual)
}
