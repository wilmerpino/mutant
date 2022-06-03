package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/wilmerpino/mutant/domain/model"
	"github.com/wilmerpino/mutant/interface/constants"
	"github.com/wilmerpino/mutant/interface/presenter"
)

func (cc *mutantController) getStats(ctx presenter.IContext) {
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

// @Summary Stats for DNA mutants vs. humans
// @Description Obtains the DNA statistics of mutants vs. humans.
// @Tags Mutant
// @Produce json
// @Success 200 {object} model.Stats "Stats for DNA strand"
// @Failure 500 {object} serverErrorResponse "Errors found"
// @Router /stats [get]
func (cc *mutantController) GetStats(ctx iris.Context) {
	cc.getStats(ctx)
}
