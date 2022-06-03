package controller

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/wilmerpino/mutant/domain/model"
	"github.com/wilmerpino/mutant/interface/constants"
	"github.com/wilmerpino/mutant/tests/mocks"
)

func TestIsMutant(t *testing.T) {
	mockContext := &mocks.MockContext{}
	mockInteractor := &mocks.MockMutantInteractor{}

	input := model.InputDNA{
		DNA: []string{
			"ATGCGA",
			"CAGTGC",
			"TTATGT",
			"AGAAGG",
			"CCCCTA",
			"TCACTG",
		},
	}

	info := model.DnaInfo{
		IsMutant: true,
		DNA:      input.DNA,
	}

	expect := model.Response{
		Message: constants.DNA_MUTANT,
	}

	dataJSON, _ := json.Marshal(input)
	mockContext.SetDNAInputJSON(string(dataJSON))

	mockContext.On("ReadJSON", &model.InputDNA{}).Return(nil)
	mockInteractor.On("SaveStrand", info).Return(nil)
	mockContext.On("JSON", expect).Return(0, nil)
	controller := &mutantController{
		mutantInteractor: mockInteractor,
	}

	controller.isMutant(mockContext)

	mockContext.AssertExpectations(t)
	mockInteractor.AssertExpectations(t)
}

func TestIsHuman(t *testing.T) {
	mockContext := &mocks.MockContext{}
	mockInteractor := &mocks.MockMutantInteractor{}

	input := model.InputDNA{
		DNA: []string{
			"ATGCGA",
			"CAGGTA",
			"TACGGT",
			"ACATGC",
			"CCACTA",
			"TCCACT"},
	}

	info := model.DnaInfo{
		IsMutant: false,
		DNA:      input.DNA,
	}

	expect := model.Response{
		Message: constants.DNA_HUMAN,
	}

	dataJSON, _ := json.Marshal(input)
	mockContext.SetDNAInputJSON(string(dataJSON))

	mockContext.On("ReadJSON", &model.InputDNA{}).Return(nil)
	mockInteractor.On("SaveStrand", info).Return(nil)
	mockContext.On("JSON", expect).Return(0, nil)
	mockContext.On("StatusCode", iris.StatusForbidden).Return()
	controller := &mutantController{
		mutantInteractor: mockInteractor,
	}

	controller.isMutant(mockContext)

	mockContext.AssertExpectations(t)
	mockInteractor.AssertExpectations(t)
}

func TestInputEmpty(t *testing.T) {
	mockContext := &mocks.MockContext{}
	mockInteractor := &mocks.MockMutantInteractor{}

	expect := model.Response{
		Message: constants.EMPTY_INPUT,
	}

	mockContext.SetDNAInputJSON("")

	mockContext.On("ReadJSON", &model.InputDNA{}).Return(nil)
	mockContext.On("Application").Return(nil)
	mockContext.On("StatusCode", iris.StatusBadRequest).Return()
	mockContext.On("JSON", expect).Return(0, nil)
	controller := &mutantController{
		mutantInteractor: mockInteractor,
	}

	controller.isMutant(mockContext)

	mockContext.AssertExpectations(t)
	mockInteractor.AssertExpectations(t)
}

func TestInputError(t *testing.T) {
	mockContext := &mocks.MockContext{}
	mockInteractor := &mocks.MockMutantInteractor{}

	expect := model.Response{
		Message: constants.READ_INPUT_ERROR,
	}

	mockContext.SetDNAInputJSON("")

	mockContext.On("ReadJSON", &model.InputDNA{}).Return(errors.New("any error"))
	mockContext.On("Application").Return(nil)
	mockContext.On("StatusCode", iris.StatusBadRequest).Return()
	mockContext.On("JSON", expect).Return(0, nil)
	controller := &mutantController{
		mutantInteractor: mockInteractor,
	}

	controller.isMutant(mockContext)

	mockContext.AssertExpectations(t)
	mockInteractor.AssertExpectations(t)
}

func TestInputArrayLenghtError(t *testing.T) {
	mockContext := &mocks.MockContext{}
	mockInteractor := &mocks.MockMutantInteractor{}

	input := model.InputDNA{
		DNA: []string{
			"ATG",
			"CAG",
			"TAC",
		},
	}

	dataJSON, _ := json.Marshal(input)
	mockContext.SetDNAInputJSON(string(dataJSON))

	expect := model.Response{
		Message: constants.STRAND_LENGTH_INVALID,
	}

	mockContext.SetDNAInputJSON(string(dataJSON))

	mockContext.On("ReadJSON", &model.InputDNA{}).Return(nil)
	mockContext.On("Application").Return(nil)
	mockContext.On("StatusCode", iris.StatusBadRequest).Return()
	mockContext.On("JSON", expect).Return(0, nil)
	controller := &mutantController{
		mutantInteractor: mockInteractor,
	}

	controller.isMutant(mockContext)

	mockContext.AssertExpectations(t)
	mockInteractor.AssertExpectations(t)
}

func TestSaveStrandError(t *testing.T) {
	mockContext := &mocks.MockContext{}
	mockInteractor := &mocks.MockMutantInteractor{}

	input := model.InputDNA{
		DNA: []string{
			"ATGCGA",
			"CAGGTA",
			"TACGGT",
			"ACATGC",
			"CCACTA",
			"TCCACT",
		},
	}

	info := model.DnaInfo{
		IsMutant: false,
		DNA:      input.DNA,
	}

	expect := model.Response{
		Message: "any error",
	}

	dataJSON, _ := json.Marshal(input)
	mockContext.SetDNAInputJSON(string(dataJSON))

	mockContext.On("ReadJSON", &model.InputDNA{}).Return(nil)
	mockInteractor.On("SaveStrand", info).Return(errors.New("any error"))
	mockContext.On("StatusCode", iris.StatusBadRequest).Return()
	mockContext.On("JSON", expect).Return(0, nil)
	controller := &mutantController{
		mutantInteractor: mockInteractor,
	}

	controller.isMutant(mockContext)

	mockContext.AssertExpectations(t)
	mockInteractor.AssertExpectations(t)
}
