package controller

import (
	"errors"
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/wilmerpino/mutant/domain/model"
	"github.com/wilmerpino/mutant/interface/constants"
	"github.com/wilmerpino/mutant/tests/mocks"
)

func TestGetStatsOk(t *testing.T) {
	mockContext := &mocks.MockContext{}
	mockInteractor := &mocks.MockMutantInteractor{}

	expect := model.Stats{}
	mockInteractor.On("Stats").Return(expect, nil)
	mockContext.On("StatusCode", iris.StatusOK).Return()
	mockContext.On("JSON", expect).Return(0, nil)

	controller := &mutantController{
		mutantInteractor: mockInteractor,
	}
	controller.getStats(mockContext)

	mockContext.AssertExpectations(t)
	mockInteractor.AssertExpectations(t)
}

func TestGetStatsError(t *testing.T) {
	mockContext := &mocks.MockContext{}
	mockInteractor := &mocks.MockMutantInteractor{}

	expect := model.Response{
		Message: constants.INTERNAL_SERVER_ERROR,
	}
	mockInteractor.On("Stats").Return(model.Stats{}, errors.New("any error"))
	mockContext.On("StatusCode", iris.StatusInternalServerError).Return()
	mockContext.On("JSON", expect).Return(0, nil)

	controller := &mutantController{
		mutantInteractor: mockInteractor,
	}
	controller.getStats(mockContext)

	mockContext.AssertExpectations(t)
	mockInteractor.AssertExpectations(t)
}
