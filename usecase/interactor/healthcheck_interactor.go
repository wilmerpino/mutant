package interactor

import (
	"github.com/wilmerpino/mutant/domain/model"
	"github.com/wilmerpino/mutant/usecase/presenter"
)

type healtCheckInteractor struct {
	HealtCheckPresenter presenter.IHealtCheckPresenter
}

type IHealtCheckInteractor interface {
	GetHealthCheck() (model.HealthCheckResponse, error)
}

func NewHealtCheckInteractor(p presenter.IHealtCheckPresenter) IHealtCheckInteractor {
	return &healtCheckInteractor{p}
}

func (h *healtCheckInteractor) GetHealthCheck() (model.HealthCheckResponse, error) {
	return h.HealtCheckPresenter.ResponseHealthCheck(), nil
}
