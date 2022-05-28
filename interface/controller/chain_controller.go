package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/wilmerpino/mutant/domain/model"
	"github.com/wilmerpino/mutant/interface/constants"
	"github.com/wilmerpino/mutant/usecase/interactor"
)

type chainController struct {
	chainInteractor interactor.IChainInteractor
}

type IChainController interface {
	GetStats(ctx iris.Context)
}

func NewChainController(c interactor.IChainInteractor) IChainController {
	return &chainController{c}
}

func (cc chainController) GetStats(ctx iris.Context) {
	dc, err := cc.chainInteractor.Stats()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ResponseError{
			Message: constants.INTERNAL_SERVER_ERROR,
			Code:    iris.StatusInternalServerError,
		})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(dc)
}
