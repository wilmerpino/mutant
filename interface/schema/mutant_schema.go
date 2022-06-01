package schema

import (
	"errors"
	"regexp"

	"github.com/kataras/iris/v12"
	"github.com/wilmerpino/mutant/domain/model"
	"github.com/wilmerpino/mutant/interface/constants"
)

func ValidateInputDNA(ctx iris.Context) ([][]string, []string, error) {
	var input model.InputDNA
	err := ctx.ReadJSON(&input)
	strand := input.DNA
	// Validates input data
	if err != nil {
		ctx.Application().Logger().Errorf(err.Error())
		return nil, strand, errors.New(constants.READ_INPUT_ERROR)
	}
	// Validates strand length
	if len(input.DNA) == 0 {
		ctx.Application().Logger().Errorf("DNA input is empty")
		return nil, strand, errors.New(constants.EMPTY_INPUT)
	}
	// Validates strand length must be >= 4
	if len(input.DNA) < 4 {
		ctx.Application().Logger().Errorf("DNA strand length must be >= 4")
		return nil, strand, errors.New(constants.STRAND_LENGTH_INVALID)
	}

	ldna := len(strand)
	dna := make([][]string, ldna, ldna)
	for i, s := range strand {
		// Validates string length
		if len(s) != ldna {
			ctx.Application().Logger().Errorf("DNA strand length not valid")
			return nil, strand, errors.New(constants.STRAND_LENGTH_INVALID)
		}
		// Validates allowed characters
		var result = regexp.MustCompile(constants.LETTERS_ALLOWED)
		if !result.MatchString(s) {
			ctx.Application().Logger().Errorf("String has some invalid character, it is allowed to %s", constants.CHARACTER_NOT_VALID)
			return nil, strand, errors.New(constants.CHARACTER_NOT_VALID)
		}
		// Load an array with the DNA strand
		dna[i] = make([]string, ldna)
		for j := 0; j < ldna; j++ {
			dna[i][j] = string(s[j])
		}
	}
	return dna, strand, nil
}

func initializeMap() map[string]int {
	return map[string]int{
		"A": 0,
		"C": 0,
		"G": 0,
		"T": 0,
	}
}

func validaSequence(result map[string]int) bool {
	for _, val := range result {
		if val >= 4 {
			return true
		}
	}

	return false
}

func ValidateStrandDNA(dna [][]string) bool {
	ldna := len(dna)
	var j int
	// Horizontal validation
	result := initializeMap()
	for i := 0; i < ldna; i++ {
		for j = 0; j < ldna; j++ {
			result[dna[i][j]]++
		}
		if validaSequence(result) {
			return true
		}
		result = initializeMap()
	}

	// Vertical validation
	result = initializeMap()
	for i := 0; i < ldna; i++ {
		for j = 0; j < ldna; j++ {
			result[dna[j][i]]++
		}
		if validaSequence(result) {
			return true
		}
		result = initializeMap()
	}

	// Oblique validation from right to left
	result = initializeMap()
	result2 := initializeMap() // new result for inverse oblique
	i := ldna - 3
	for k := 0; k < ldna; k++ {
		i = ldna - 3 + k
		for j = ldna - 1; j >= 0; j-- {
			result[dna[i][j]]++
			result2[dna[j][i]]++ // check inverse oblique
			if i == 0 {
				break
			}
			i -= 1
		}
		if validaSequence(result) || validaSequence(result2) {
			return true
		}
		if i == j { // breaks when the main diagonal is checked
			break
		}
		result = initializeMap()
		result2 = initializeMap()
	}

	// Oblique validation from left to right
	i = ldna - 3
	n := 0
	j = 0
	m := i
	result = initializeMap()
	for k := 0; k < ldna-1; k++ {
		if m < ldna {
			i = ldna - 3 + k
			m++
		} else {
			i = ldna - 1
			n = m - i + k - 3
		}
		for j = n; j < m; j++ {
			result[dna[i][j]]++
			i--
		}
		if validaSequence(result) {
			return true
		}
		result = initializeMap()
	}

	return false
}
