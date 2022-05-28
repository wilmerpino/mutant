package schema

import (
	"errors"

	"github.com/kataras/iris/v12"
	"github.com/wilmerpino/mutant/domain/model"
	"github.com/wilmerpino/mutant/interface/constants"
)

func ValidateInputDNA(ctx iris.Context) ([]string, error) {
	var input model.InputDNA
	err := ctx.ReadJSON(&input)
	if err != nil {
		ctx.Application().Logger().Errorf(err.Error())
		return nil, errors.New(constants.READ_INPUT_ERROR)
	}

	if len(input.DNA) == 0 {
		ctx.Application().Logger().Errorf("DNA input is empty")
		return nil, errors.New(constants.EMPTY_INPUT)
	}

	dna := input.DNA
	ldna := len(dna)
	for _, s := range dna {
		if len(s) != ldna {
			ctx.Application().Logger().Errorf("DNA strand length not valid")
			return nil, errors.New(constants.STRAND_LENGTH_INVALID)
		}
	}

	return dna, nil
}
