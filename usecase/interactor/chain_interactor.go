package interactor

import (
	"github.com/wilmerpino/mutant/domain/model"
	"github.com/wilmerpino/mutant/usecase/presenter"
	"github.com/wilmerpino/mutant/usecase/repository"
)

type chainInteractor struct {
	ChainRepository repository.IChainRepository
	ChainPresenter  presenter.IChainPresenter
}

type IChainInteractor interface {
	Stats() (model.Stats, error)
}

func NewChainInteractor(r repository.IChainRepository, p presenter.IChainPresenter) IChainInteractor {
	return &chainInteractor{r, p}
}

func (dc *chainInteractor) Stats() (model.Stats, error) {
	m, err := dc.ChainRepository.FindAll()
	if err != nil {
		return model.Stats{}, err
	}

	return dc.ChainPresenter.ResponseChainsStats(m), nil
}
