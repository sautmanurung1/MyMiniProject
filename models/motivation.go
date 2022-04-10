package models

import "gorm.io/gorm"

type Motivation struct {
	gorm.Model

	ID         int    `json:"id" form:"id" gorm:"primaryKey"`
	Motivation string `json:"motivation" form:"motivation"`
}
