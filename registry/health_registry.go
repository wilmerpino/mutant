package registry

import (
	"github.com/wilmerpino/mutant/interface/controller"
	ip "github.com/wilmerpino/mutant/interface/presenter"
	"github.com/wilmerpino/mutant/usecase/interactor"
	up "github.com/wilmerpino/mutant/usecase/presenter"
)

func (r *registry) NewHealtCheckController() controller.IHealtCheckController {
	return controller.NewHealtCheckController(r.NewHealtCheckInteractor())
}

func (r *registry) NewHealtCheckInteractor() interactor.IHealtCheckInteractor {
	return interactor.NewHealtCheckInteractor(r.NewHealtCheckPresenter())
}

func (r *registry) NewHealtCheckPresenter() up.IHealtCheckPresenter {
	return ip.NewHealtCheckPresenter()
}
