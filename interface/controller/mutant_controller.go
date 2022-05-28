package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/wilmerpino/mutant/domain/model"
	"github.com/wilmerpino/mutant/interface/schema"
)

func (cc *mutantController) IsMutant(ctx iris.Context) {
	dna, err := schema.ValidateInputDNA(ctx)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(dna)
}
