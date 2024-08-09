package graph

import "gorm.io/gorm"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *gorm.DB // 데이터베이스 연결 객체 추가
}
