package voucher

import (
	"gorm.io/gorm"
	"time"
)

type Voucher struct {
	gorm.Model
	Code           string    `gorm:"not null; default null; unique"`
	Comment        string    `gorm:"not null; default null"`
	Start          time.Time `gorm:"not null; default null; type:date"`
	End            time.Time `gorm:"default null; type:date"`
	DurationInDays int       `gorm:"default null"`
	IsActive       bool      `gorm:"not null; default false"`
	MultiUse       bool      `gorm:"not null; default false"`
}
