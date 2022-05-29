package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/wilmerpino/mutant/domain/model"
	"github.com/wilmerpino/mutant/interface/constants"
)

func (cc *mutantController) GetStats(ctx iris.Context) {
	dc, err := cc.mutantInteractor.Stats()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.Response{
			Message: constants.INTERNAL_SERVER_ERROR,
		})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(dc)
}
