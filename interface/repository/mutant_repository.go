package repository

import (
	"strings"
	"time"

	"github.com/wilmerpino/mutant/domain/model"
	"gorm.io/gorm"
)

type mutantRepository struct {
	db *gorm.DB
}

type IMutantRepository interface {
	Save(Mutant model.DnaInfo) error
	FindAll() ([]model.Strand, error)
}

func NewMutantRepository(db *gorm.DB) IMutantRepository {
	return &mutantRepository{db}
}

func (cr *mutantRepository) Save(data model.DnaInfo) error {
	strant := model.Strand{
		Strand:   strings.Join(data.DNA, " "),
		IsMutant: data.IsMutant,
		Date:     time.Now(),
	}

	result := cr.db.Create(&strant)
	return result.Error
}

func (cr *mutantRepository) FindAll() ([]model.Strand, error) {
	var strant []model.Strand
	result := cr.db.Find(&strant)
	return strant, result.Error
}
