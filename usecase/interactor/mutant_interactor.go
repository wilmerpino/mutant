package interactor

import (
	"github.com/wilmerpino/mutant/domain/model"
	"github.com/wilmerpino/mutant/usecase/presenter"
	"github.com/wilmerpino/mutant/usecase/repository"
)

type mutantInteractor struct {
	MutantRepository repository.IMutantRepository
	MutantPresenter  presenter.IMutantPresenter
}

type IMutantInteractor interface {
	Stats() (model.Stats, error)
	SaveStrand(model.DnaInfo) error
}

func NewMutantInteractor(r repository.IMutantRepository, p presenter.IMutantPresenter) IMutantInteractor {
	return &mutantInteractor{r, p}
}

func (dc *mutantInteractor) Stats() (model.Stats, error) {
	m, err := dc.MutantRepository.FindAll()
	if err != nil {
		return model.Stats{}, err
	}

	return dc.MutantPresenter.ResponseMutantsStats(m), nil
}

func (dc *mutantInteractor) SaveStrand(data model.DnaInfo) error {
	return dc.MutantRepository.Save(data)
}
