package schema

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/wilmerpino/mutant/domain/model"
	"github.com/wilmerpino/mutant/interface/constants"
	"github.com/wilmerpino/mutant/interface/presenter"
)

func ValidateInputDNA(ctx presenter.IContext) ([][]string, []string, error) {
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

func ValidateStrandDNA(dna [][]string) bool {
	ldna := len(dna)
	var j int
	// Horizontal validation
	var last string
	for i := 0; i < ldna; i++ {
		result := 0
		last = ""
		for j = 0; j < ldna; j++ {
			if last == dna[i][j] {
				result++
				if result >= 3 {
					fmt.Printf("1- LETRA: %s, I: %d - J: %d\n", dna[i][j], i, j)
					return true
				}
			} else {
				result = 0
			}
			last = dna[i][j]
		}
	}

	// Vertical validation
	var result int
	for i := 0; i < ldna; i++ {
		result = 0
		last = ""
		for j = 0; j < ldna; j++ {
			if last == dna[j][i] {
				result++
				if result >= 3 {
					fmt.Printf("2- LETRA: %s, I: %d - J: %d\n", dna[j][j], j, i)
					return true
				}
			} else {
				result = 0
			}
			last = dna[j][i]
		}
	}

	// Principal left to right oblique validation
	result = 0
	last = ""
	for i := 0; i < ldna; i++ {
		if dna[i][i] == last {
			result++
			if result >= 3 {
				fmt.Printf("3- LETRA: %s, I: %d - J: %d\n", dna[i][i], i, i)
				return true
			}
		} else {
			result = 0
		}
		last = dna[i][i]
	}
	// Principal rigth to left oblique validation
	result = 0
	last = ""
	j = 0
	for i := ldna - 1; i >= 0; i-- {
		if dna[i][j] == last {
			result++
			if result >= 3 {
				fmt.Printf("3- LETRA: %s, I: %d - J: %d\n", dna[i][i], i, i)
				return true
			}
		} else {
			result = 0
		}
		last = dna[i][j]
		j++
	}
	// Oblique validation from right to left
	var last2 string
	var result2 int
	for k := 3; k < ldna-1; k++ {
		i := k
		result = 0
		result2 = 0
		last = ""
		last2 = ""
		for j = ldna - 1; j >= 0; j-- {
			if dna[i][j] == last {
				result++
				if result >= 3 {
					fmt.Printf("4- LETRA: %s, I: %d - J: %d\n", dna[i][j], i, j)
					return true
				}
			} else {
				result = 0
			}
			if dna[j][i] == last2 { // check inverse oblique
				result2++
				if result2 >= 3 {
					fmt.Printf("5- LETRA: %s, I: %d - J: %d\n", dna[j][i], j, i)
					return true
				}
			} else {
				result2 = 0
			}
			if i == 0 {
				break
			}
			last = dna[i][j]
			last2 = dna[j][i]
			i -= 1
		}
	}

	// Oblique validation from left to right
	i := 3
	n := 0
	j = 0
	m := i
	for k := 0; k < ldna-1; k++ {
		result = 0
		last = ""
		for j = n; j < m; j++ {
			fmt.Printf("I: %d - J: %d\n", i, j)
			if dna[i][j] == last {
				result++
				if result >= 3 {
					return true
				}
			} else {
				result = 0
			}
			last = dna[i][j]
			i--
		}
		if m < ldna-1 {
			m++
			i = m
		} else {
			i = ldna - 1
			n++
		}
		fmt.Println()
	}
	return false
}
