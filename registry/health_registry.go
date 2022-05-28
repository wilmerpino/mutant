package registry

import (
	"github.com/wilmerpino/mutant/interface/controller"
	ip "github.com/wilmerpino/mutant/interface/presenter"
	"github.com/wilmerpino/mutant/usecase/interactor"
	up "github.com/wilmerpino/mutant/usecase/presenter"
)

func (r *registry) NewHealthCheckController() controller.IHealthCheckController {
	return controller.NewHealthCheckController(r.NewHealthCheckInteractor())
}

func (r *registry) NewHealthCheckInteractor() interactor.IHealthCheckInteractor {
	return interactor.NewHealthCheckInteractor(r.NewHealthCheckPresenter())
}

func (r *registry) NewHealthCheckPresenter() up.IHealthCheckPresenter {
	return ip.NewHealthCheckPresenter()
}
