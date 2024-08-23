package db_model

import (
	"fmt"
	"go_test/graph/gql_model"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model // GORM의 내장 모델을 포함시킵니다.
	Name       string
	Email      string `gorm:"type:varchar(100);unique_index"`
	Posts      []Post `gorm:"foreignKey:AuthorID"` // User는 여러 Post를 가질 수 있음 (One-to-Many)
}

// 테이블 이름 지정
func (User) TableName() string {
	return "users_table"
}

func (u *User) ToGraphQLModel() *gql_model.User {
	createdAt := u.CreatedAt.Format(time.RFC3339)
	updatedAt := u.UpdatedAt.Format(time.RFC3339)

	// Posts를 GraphQL 모델로 변환
	var posts []*gql_model.Post
	for _, post := range u.Posts {
		posts = append(posts, post.ToGraphQLModel()) // Post도 변환
	}

	return &gql_model.User{
		ID:        fmt.Sprintf("%d", u.ID),
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		Posts:     posts, // Posts 필드 추가
	}
}
