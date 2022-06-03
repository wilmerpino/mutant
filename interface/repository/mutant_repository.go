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
	CountAll() []model.StatsResult
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

func (cr *mutantRepository) CountAll() []model.StatsResult {
	var results []model.StatsResult
	cr.db.Table("strands").
		Select("COUNT(*) as cant, is_mutant").
		Group("is_mutant").
		Scan(&results)

	return results
}
