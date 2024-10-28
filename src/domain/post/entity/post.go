package entity

import (
	"fmt"
	"go_test/graph/gql_model"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model // GORM의 내장 모델을 포함시킵니다.
	Title      string	`gorm:"type:varchar(255)"`
	Content    string	`gorm:"type:text"`
	Author     string	`gorm:"type:varchar(100)"`
}

// 테이블 이름 지정
func (Post) TableName() string {
	return "posts_table"
}

// ToGraphQLModel 메서드
func (p *Post) ToGraphQLModel() *gql_model.Post {
	createdAt := p.CreatedAt.Format(time.RFC3339)
	updatedAt := p.UpdatedAt.Format(time.RFC3339)
	return &gql_model.Post{
		ID:        fmt.Sprintf("%d", p.ID),
		Title:     p.Title,
		Content:   p.Content,
		Author:    p.Author,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	}
}
