package controller

import (
	"errors"
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/wilmerpino/mutant/domain/model"
	"github.com/wilmerpino/mutant/tests/mocks"
)

func TestGetHealthCheck(t *testing.T) {
	mockHealthInteractor := &mocks.MockHealthCheckInteractor{}
	mockContext := &mocks.MockContext{}

	expect := model.HealthCheckResponse{}
	mockHealthInteractor.On("GetHealthCheck").Return(expect, nil)
	mockContext.On("JSON", expect).Return(0, nil)

	controller := &healthCheckController{
		healthCheckInteractor: mockHealthInteractor,
	}

	controller.getHealthCheck(mockContext)

	mockContext.AssertExpectations(t)
	mockHealthInteractor.AssertExpectations(t)
}

func TestGetHealthCheckError(t *testing.T) {
	mockHealthInteractor := &mocks.MockHealthCheckInteractor{}
	mockContext := &mocks.MockContext{}

	expect := model.HealthCheckResponse{}
	mockHealthInteractor.On("GetHealthCheck").Return(expect, errors.New("any error"))
	mockContext.On("StatusCode", iris.StatusInternalServerError).Return()
	mockContext.On("StopExecution").Return()

	controller := &healthCheckController{
		healthCheckInteractor: mockHealthInteractor,
	}

	controller.getHealthCheck(mockContext)

	mockContext.AssertExpectations(t)
	mockHealthInteractor.AssertExpectations(t)
}
