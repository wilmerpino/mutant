package repository

import (
	"github.com/wilmerpino/mutant/domain/model"
)

type IMutantRepository interface {
	FindAll() ([]model.Strand, error)
	Save(model.DnaInfo) error
}
