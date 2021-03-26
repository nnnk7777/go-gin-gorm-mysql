package model

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primary_key;not null;autoIncrement:true"`
	Name      string    `gorm:"type:text;not null"`
	Mail      string    `gorm:"type:text;not null"`
	Timestamp time.Time `sql:"DEFAULT:current_timestamp"`
}
