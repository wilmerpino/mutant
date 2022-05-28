package registry

import (
	"github.com/wilmerpino/mutant/interface/controller"
	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type IRegistry interface {
	NewAppController() controller.AppController
	NewCheckController() controller.CheckController
}

func NewRegistry(db *gorm.DB) IRegistry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewChainController()
}

func (r *registry) NewCheckController() controller.CheckController {
	return r.NewHealtCheckController()
}
