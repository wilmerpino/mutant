package presenter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilmerpino/mutant/domain/model"
)

func TestResponseMutantsStats(t *testing.T) {
	presenter := &mutantPresenter{}

	expect := model.Stats{}
	var ac []model.StatsResult
	actual := presenter.ResponseMutantsStats(ac)
	assert.Equal(t, expect, actual)
}
