package interactor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilmerpino/mutant/domain/model"
	"github.com/wilmerpino/mutant/tests/mocks"
)

func TestNewServiceInteractor(t *testing.T) {
	mServicePresenter := &mocks.MockHealthCheckPresenter{}

	expected := healthCheckInteractor{
		HealthCheckPresenter: mServicePresenter,
	}

	result := NewHealthCheckInteractor(mServicePresenter)

	assert.Equal(t, &expected, result)
	mServicePresenter.AssertExpectations(t)
}

func TestHealthcheckOK(t *testing.T) {
	mServicePresenter := &mocks.MockHealthCheckPresenter{}
	expected := model.HealthCheckResponse{
		Status: "OK",
	}
	mServicePresenter.On(
		"ResponseHealthCheck",
	).Return(expected)
	interactor := healthCheckInteractor{
		HealthCheckPresenter: mServicePresenter,
	}
	obj, err := interactor.GetHealthCheck()
	assert.NoError(t, err)
	assert.Equal(t, expected, obj)

	mServicePresenter.AssertExpectations(t)
}
