package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/wilmerpino/mutant/domain/model"
	"github.com/wilmerpino/mutant/interface/constants"
	"github.com/wilmerpino/mutant/interface/schema"
)

func (cc *mutantController) IsMutant(ctx iris.Context) {
	dna, strand, err := schema.ValidateInputDNA(ctx)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.Response{
			Message: err.Error(),
		})
		return
	}

	isMutant := schema.ValidateStrandDNA(dna)
	info := model.DnaInfo{
		IsMutant: isMutant,
		DNA:      strand,
	}

	err = cc.mutantInteractor.SaveStrand(info)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.Response{
			Message: err.Error(),
		})
	}

	if isMutant {
		ctx.JSON(model.Response{
			Message: constants.DNA_MUTANT,
		})
		return
	}

	ctx.StatusCode(iris.StatusForbidden)
	ctx.JSON(model.Response{
		Message: constants.DNA_HUMAN,
	})
}
