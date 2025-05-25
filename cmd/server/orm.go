package main

import (
	// "github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Applicant struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255"`
	Birth       datatypes.Date
	Email       string   `gorm:"uniqueIndex"`
	CPF         string   `gorm:"size:11"`
	Professions []string `gorm:"type:text[]"`
}
