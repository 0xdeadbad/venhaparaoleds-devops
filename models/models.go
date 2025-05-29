package models

import (
	"time"

	"gorm.io/gorm"
)

type Applicant struct {
	gorm.Model
	Name        string       `gorm:"size:50" json:"name"`
	Birth       time.Time    `json:"birth"`
	Email       string       `gorm:"uniqueIndex" json:"email"`
	CPF         string       `gorm:"size:11;primaryKey" json:"cpf"`
	Professions []Profession `gorm:"foreignKey:ApplicantID" json:"professions"`
}

type Concourse struct {
	gorm.Model
	ID       uint
	Org      string `json:"org"`
	Edital   string `json:"edital"`
	ConcCode uint64 `gorm:"primaryKey;uniqueIndex" json:"conc_code"`
}

type Vacancy struct {
	gorm.Model
	ConcID int       `gorm:"uniqueIndex"`
	Conc   Concourse `gorm:"foreignKey:ConcID"`
	Name   string    `gorm:"uniqueIndex" json:"name"`
}

type Profession struct {
	gorm.Model
	ApplicantID uint
	Name        string `json:"name"`
}
