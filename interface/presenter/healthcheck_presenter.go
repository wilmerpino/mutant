package presenter

import (
	"github.com/wilmerpino/mutant/domain/model"
)

type healthCheckPresenter struct{}

type IHealthCheckPresenter interface {
	ResponseHealthCheck() model.HealthCheckResponse
}

func NewHealthCheckPresenter() IHealthCheckPresenter {
	return &healthCheckPresenter{}
}

func (h *healthCheckPresenter) ResponseHealthCheck() model.HealthCheckResponse {
	return model.HealthCheckResponse{
		Status: "OK",
	}
}
