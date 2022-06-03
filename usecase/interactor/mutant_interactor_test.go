package interactor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilmerpino/mutant/domain/model"
	"github.com/wilmerpino/mutant/tests/mocks"
)

func TestNewMutantInteractor(t *testing.T) {
	mockRepository := &mocks.MockMutantRepository{}
	mockPresenter := &mocks.MockMutantPresenter{}

	expected := &mutantInteractor{
		mockRepository,
		mockPresenter,
	}

	actual := NewMutantInteractor(
		mockRepository,
		mockPresenter,
	)

	assert.Equal(t, expected, actual)
	mockRepository.AssertExpectations(t)
	mockPresenter.AssertExpectations(t)
}

func TestStatsOK(t *testing.T) {
	mockRepository := &mocks.MockMutantRepository{}
	mockPresenter := &mocks.MockMutantPresenter{}

	expected := model.Stats{}

	mockRepository.On("CountAll").Return([]model.StatsResult{})
	mockPresenter.On("ResponseMutantsStats", []model.StatsResult{}).Return(expected)

	interactor := NewMutantInteractor(
		mockRepository,
		mockPresenter,
	)

	actual, err := interactor.Stats()

	assert.Equal(t, expected, actual)
	assert.Nil(t, err)
	mockRepository.AssertExpectations(t)
	mockPresenter.AssertExpectations(t)
}

func TestSaveStrandOK(t *testing.T) {
	mockRepository := &mocks.MockMutantRepository{}
	mockPresenter := &mocks.MockMutantPresenter{}

	data := model.DnaInfo{
		DNA:      []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
		IsMutant: true,
	}

	mockRepository.On("Save", data).Return(nil)

	interactor := NewMutantInteractor(
		mockRepository,
		mockPresenter,
	)

	err := interactor.SaveStrand(data)

	assert.Nil(t, err)
	mockRepository.AssertExpectations(t)
	mockPresenter.AssertExpectations(t)
}
