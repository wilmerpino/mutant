package repository

import (
	"github.com/wilmerpino/mutant/domain/model"
)

type IChainRepository interface {
	FindAll() ([]*model.DnaChain, error)
}
