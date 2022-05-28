package presenter

import (
	"github.com/wilmerpino/mutant/domain/model"
)

type healtCheckPresenter struct{}

type IHealtCheckPresenter interface {
	ResponseHealthCheck() model.HealthCheckResponse
}

func NewHealtCheckPresenter() IHealtCheckPresenter {
	return &healtCheckPresenter{}
}

func (h *healtCheckPresenter) ResponseHealthCheck() model.HealthCheckResponse {
	return model.HealthCheckResponse{
		Status: "OK",
	}
}
