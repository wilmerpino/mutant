package registry

import (
	"github.com/wilmerpino/mutant/interface/controller"
	ip "github.com/wilmerpino/mutant/interface/presenter"
	ir "github.com/wilmerpino/mutant/interface/repository"
	"github.com/wilmerpino/mutant/usecase/interactor"
	up "github.com/wilmerpino/mutant/usecase/presenter"
	ur "github.com/wilmerpino/mutant/usecase/repository"
)

func (r *registry) NewChainController() controller.IChainController {
	return controller.NewChainController(r.NewChainInteractor())
}

func (r *registry) NewChainInteractor() interactor.IChainInteractor {
	return interactor.NewChainInteractor(r.NewChainRepository(), r.NewChainPresenter())
}

func (r *registry) NewChainRepository() ur.IChainRepository {
	return ir.NewChainRepository(r.db)
}

func (r *registry) NewChainPresenter() up.IChainPresenter {
	return ip.NewChainPresenter()
}
