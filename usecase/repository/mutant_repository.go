package repository

import (
	"github.com/wilmerpino/mutant/domain/model"
)

type IMutantRepository interface {
	FindAll() ([]*model.DnaChain, error)
}
