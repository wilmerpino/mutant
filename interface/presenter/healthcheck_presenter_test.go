package presenter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilmerpino/mutant/domain/model"
)

func TestNewHealthCheckPresenter(t *testing.T) {
	expect := &healthCheckPresenter{}
	actual := NewHealthCheckPresenter()
	assert.Equal(t, expect, actual)
}

func TestResponseHealthCheck(t *testing.T) {
	expect := model.HealthCheckResponse{
		Status: "OK",
	}
	presenter := &healthCheckPresenter{}
	actual := presenter.ResponseHealthCheck()
	assert.Equal(t, expect, actual)
}
