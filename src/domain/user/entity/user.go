package entity

import (
	"fmt"
	"go_test/graph/gql_model"
	"go_test/src/domain/post/entity"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model // GORM의 내장 모델을 포함시킵니다.
	Name       string			`gorm:"type:varchar(100);unique"`
	Email      string			`gorm:"type:varchar(100);unique_index"`
	Posts      []entity.Post	`gorm:"foreignKey:Author;references:Name"`
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
		Posts:     posts,
	}
}
