package voucher

import (
	"gorm.io/gorm"
)

type Voucher struct {
	gorm.Model
	Code           string `gorm:"not null; default null; unique"`
	Comment        string `gorm:"not null; default null"`
	Start          string `gorm:"not null; default null; type:date"`
	End            string `gorm:"default null; type:date"`
	DurationInDays int    `gorm:"default null"`
	IsActive       bool   `gorm:"not null; default false"`
	MultiUse       bool   `gorm:"not null; default false"`
}
