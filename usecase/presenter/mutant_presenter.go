package presenter

import (
	"github.com/wilmerpino/mutant/domain/model"
)

type IMutantPresenter interface {
	ResponseMutantsStats(u []model.Strand) model.Stats
}
