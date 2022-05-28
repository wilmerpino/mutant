package presenter

import "github.com/wilmerpino/mutant/domain/model"

type mutantPresenter struct{}

type IMutantPresenter interface {
	ResponseMutantsStats(ac []*model.DnaChain) model.Stats
}

func NewMutantPresenter() IMutantPresenter {
	return &mutantPresenter{}
}

func (cp *mutantPresenter) ResponseMutantsStats(ac []*model.DnaChain) model.Stats {
	cMutant, cHuman := 0, 0
	ratio := float32(1.0)
	for _, c := range ac {
		if c.IsMutant {
			cMutant++
		} else {
			cHuman++
		}
	}

	if cHuman != 0 {
		ratio = float32(cMutant) / float32(cHuman)
	}

	return model.Stats{
		CountMutant: int32(cMutant),
		CountHuman:  int32(cHuman),
		Ratio:       ratio,
	}
}
