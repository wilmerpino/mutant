package presenter

import (
	"github.com/wilmerpino/mutant/domain/model"
)

type mutantPresenter struct{}

type IMutantPresenter interface {
	ResponseMutantsStats(ac []model.StatsResult) model.Stats
}

func NewMutantPresenter() IMutantPresenter {
	return &mutantPresenter{}
}

func (cp *mutantPresenter) ResponseMutantsStats(count []model.StatsResult) model.Stats {
	cMutant, cHuman := 0, 0
	ratio := float32(0)

	if len(count) > 0 {
		if count[0].IsMutant {
			cMutant = count[0].Cant
		}
		if len(count) > 1 && count[1].IsMutant {
			cMutant = count[1].Cant
		}
		if !count[0].IsMutant {
			cHuman = count[0].Cant
		}
		if len(count) > 1 && !count[1].IsMutant {
			cHuman = count[1].Cant
		}
		if cHuman != 0 {
			ratio = float32(cMutant) / float32(cHuman)
		}
	}

	return model.Stats{
		CountMutant: int32(cMutant),
		CountHuman:  int32(cHuman),
		Ratio:       ratio,
	}
}
