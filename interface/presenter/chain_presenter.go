package presenter

import "github.com/wilmerpino/mutant/domain/model"

type chainPresenter struct{}

type IChainPresenter interface {
	ResponseChainsStats(ac []*model.DnaChain) model.Stats
}

func NewChainPresenter() IChainPresenter {
	return &chainPresenter{}
}

func (cp chainPresenter) ResponseChainsStats(ac []*model.DnaChain) model.Stats {
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
