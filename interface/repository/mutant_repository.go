package repository

import (
	"github.com/wilmerpino/mutant/domain/model"
	"gorm.io/gorm"
)

type mutantRepository struct {
	db *gorm.DB
}

type IMutantRepository interface {
	Save(Mutant model.DnaInfo) error
	FindAll() ([]*model.DnaChain, error)
}

func NewMutantRepository(db *gorm.DB) IMutantRepository {
	return &mutantRepository{db}
}

func (cr *mutantRepository) Save(data model.DnaInfo) error {
	return nil
}

func (cr *mutantRepository) FindAll() ([]*model.DnaChain, error) {
	var dc []*model.DnaChain
	return dc, nil
}
