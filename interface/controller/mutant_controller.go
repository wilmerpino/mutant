package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/wilmerpino/mutant/domain/model"
	"github.com/wilmerpino/mutant/interface/constants"
	"github.com/wilmerpino/mutant/interface/presenter"
	"github.com/wilmerpino/mutant/interface/schema"
)

type mutantResponse struct {
	Message string `json:"message" example:"DNA_MUTANT"`
}

type humanResponse struct {
	Message string `json:"message" example:"DNA_HUMAN"`
}

type inputErrorResponse struct {
	Message string `json:"message" example:"STRAND_LENGTH_INVALID"`
}

type serverErrorResponse struct {
	Message string `json:"message" example:"INTERNAL_SERVER_ERROR"`
}

func (cc *mutantController) isMutant(ctx presenter.IContext) {
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
		return
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

// @Summary DNA strand is human or mutant
// @Description Validates whether the DNA strand is human or mutant
// @Tags Mutant
// @Produce json
// @Param  data body model.InputDNA true "DNA strand"
// @Success 200 {object} mutantResponse "DNA is mutant"
// @Failure 403 {object} humanResponse "DNA is human"
// @Failure 400 {object} inputErrorResponse "Input error"
// @Failure 500 {object} serverErrorResponse "Errors found"
// @Router /mutant [post]
func (cc *mutantController) IsMutant(ctx iris.Context) {
	cc.isMutant(ctx)
}
