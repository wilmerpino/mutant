package registry

import (
	"github.com/wilmerpino/mutant/interface/controller"
	ip "github.com/wilmerpino/mutant/interface/presenter"
	ir "github.com/wilmerpino/mutant/interface/repository"
	"github.com/wilmerpino/mutant/usecase/interactor"
	up "github.com/wilmerpino/mutant/usecase/presenter"
	ur "github.com/wilmerpino/mutant/usecase/repository"
)

func (r *registry) NewMutantController() controller.IMutantController {
	return controller.NewMutantController(r.NewMutantInteractor())
}

func (r *registry) NewMutantInteractor() interactor.IMutantInteractor {
	return interactor.NewMutantInteractor(r.NewMutantRepository(), r.NewMutantPresenter())
}

func (r *registry) NewMutantRepository() ur.IMutantRepository {
	return ir.NewMutantRepository(r.db)
}

func (r *registry) NewMutantPresenter() up.IMutantPresenter {
	return ip.NewMutantPresenter()
}
