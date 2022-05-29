package model

import "time"

type Strand struct {
	Id       int       `json:"id" gorm:"primaryKey"`
	Strand   string    `json:"strand"`
	IsMutant bool      `json:"is_mutant"`
	Date     time.Time `json:"date"`
}
