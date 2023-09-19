package voucher

import (
	"database/sql"
	"errors"

	"gorm.io/gorm"
)

type Voucher struct {
	gorm.Model
	Code           string         `gorm:"not null; default null; unique"`
	Comment        string         `gorm:"not null; default null"`
	Start          sql.NullString `gorm:"default null; type:date"`
	End            sql.NullString `gorm:"default null; type:date"`
	DurationInDays int            `gorm:"default null"`
	IsActive       bool           `gorm:"not null; default false"`
	MultiUse       bool           `gorm:"not null; default false"`
}

type CreateVoucherRequest struct {
	Code           string
	Comment        string
	Start          string
	End            string
	DurationInDays int
	IsActive       bool
	MultiUse       bool
	StartWhenEnter bool
}

func (c CreateVoucherRequest) ToVoucher() (Voucher, error) {
	var start sql.NullString = sql.NullString{String: c.Start, Valid: true}
	var end sql.NullString = sql.NullString{String: c.End, Valid: true}

	if c.StartWhenEnter {
		start.Valid = false
	}

	if len(c.End) == 0 {
		end.Valid = false
	}

	if c.MultiUse && c.StartWhenEnter {
		return Voucher{}, errors.New("MultiUse and StartWhenEnter can't be set at the same time")
	}

	return Voucher{
		Code:           c.Code,
		Comment:        c.Comment,
		Start:          start,
		End:            end,
		DurationInDays: c.DurationInDays,
		IsActive:       c.IsActive,
		MultiUse:       c.MultiUse,
	}, nil
}
