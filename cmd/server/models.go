package main

import (
	// "gorm.io/driver/postgres"
	"time"

	"gorm.io/gorm"
)

type Applicant struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"size:50"`
	Birth       time.Time
	Email       string       `gorm:"uniqueIndex"`
	CPF         string       `gorm:"size:11;primaryKey"`
	Professions []Profession `gorm:"type:int[];references:Profession"`
}

type Concourse struct {
	gorm.Model
	Org       string
	Edital    string
	ConcCode  uint64    `gorm:"primaryKey"`
	Vacancies []Vacancy `gorm:"type:int[];references:Vacancy"`
}

type Vacancy struct {
	gorm.Model
	Name string `gorm:"uniqueIndex" json:"name"`
}

type Profession struct {
	gorm.Model
	Name string `gorm:"uniqueIndex" json:"name"`
}
