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
	Professions []Profession `gorm:"many2many:prof_app" json:"professions"` // ``
}

type Concourse struct {
	gorm.Model
	ID       uint
	Org      string `json:"org"`
	Edital   string `json:"edital"`
	ConcCode string `gorm:"primaryKey;uniqueIndex" json:"conc_code"`
}

type Vacancy struct {
	gorm.Model
	ConcID         uint       `gorm:"uniqueIndex"`
	Conc           Concourse  `gorm:"foreignKey:ConcID"`
	ProfessionName string     `gorm:"uniqueIndex"`
	Profession     Profession `gorm:"foreignKey:ProfessionName"`
}

type Profession struct {
	gorm.Model
	ApplicantID uint
	NameSlug    string `gorm:"primaryKey,uniqueIndex" json:"name_slug"`
	Name        string `gorm:"uniqueIndex" json:"name"`
	// Applicants []*Applicant `gorm:"many2many:applicant_professions;"`
}
