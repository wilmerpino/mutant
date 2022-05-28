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
	NewHealthController() controller.HealthCheckController
}

func NewRegistry(db *gorm.DB) IRegistry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewMutantController()
}

func (r *registry) NewHealthController() controller.HealthCheckController {
	return r.NewHealthCheckController()
}
