package db_model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model // GORM의 내장 모델을 포함시킵니다.
	Name       string
	Email      string `gorm:"type:varchar(100);unique_index"`
}
