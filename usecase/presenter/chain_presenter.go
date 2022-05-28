package presenter

import (
	"github.com/wilmerpino/mutant/domain/model"
)

type IChainPresenter interface {
	ResponseChainsStats(u []*model.DnaChain) model.Stats
}
