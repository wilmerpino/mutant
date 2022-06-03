package repository

import (
	"github.com/wilmerpino/mutant/domain/model"
)

type IMutantRepository interface {
	Save(model.DnaInfo) error
	CountAll() []model.StatsResult
}
