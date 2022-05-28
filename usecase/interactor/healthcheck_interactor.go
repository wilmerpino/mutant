package interactor

import (
	"github.com/wilmerpino/mutant/domain/model"
	"github.com/wilmerpino/mutant/usecase/presenter"
)

type healthCheckInteractor struct {
	HealthCheckPresenter presenter.IHealthCheckPresenter
}

type IHealthCheckInteractor interface {
	GetHealthCheck() (model.HealthCheckResponse, error)
}

func NewHealthCheckInteractor(p presenter.IHealthCheckPresenter) IHealthCheckInteractor {
	return &healthCheckInteractor{p}
}

func (h *healthCheckInteractor) GetHealthCheck() (model.HealthCheckResponse, error) {
	return h.HealthCheckPresenter.ResponseHealthCheck(), nil
}
