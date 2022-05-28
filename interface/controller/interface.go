package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/wilmerpino/mutant/usecase/interactor"
)

type mutantController struct {
	mutantInteractor interactor.IMutantInteractor
}

type IMutantController interface {
	GetStats(ctx iris.Context)
	IsMutant(ctx iris.Context)
}

func NewMutantController(c interactor.IMutantInteractor) IMutantController {
	return &mutantController{c}
}
