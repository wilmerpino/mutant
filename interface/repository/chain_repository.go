package repository

import (
	"github.com/wilmerpino/mutant/domain/model"
	"gorm.io/gorm"
)

type chainRepository struct {
	db *gorm.DB
}

type IChainRepository interface {
	Save(chain model.DnaInfo) error
	FindAll() ([]*model.DnaChain, error)
}

func NewChainRepository(db *gorm.DB) IChainRepository {
	return &chainRepository{db}
}

func (cr *chainRepository) Save(data model.DnaInfo) error {
	return nil
}

func (cr *chainRepository) FindAll() ([]*model.DnaChain, error) {
	var dc []*model.DnaChain
	return dc, nil
}
