package model

import "time"

type Model struct {
	Id        uint64 `gorm:"primary_key"`
	CreatedBy uint64
	CreatedAt time.Time
	UpdatedBy uint64
	UpdatedAt time.Time
}
