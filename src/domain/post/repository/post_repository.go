package repository

import (
	"go_test/src/domain/post/entity"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) CreatePost(post *entity.Post) (*entity.Post, error) {
	if err := r.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (r *PostRepository) UpdatePost(post *entity.Post) (*entity.Post, error) {
	if err := r.db.Save(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (r *PostRepository) DeletePost(post *entity.Post) error {
	return r.db.Delete(post).Error
}

func (r *PostRepository) GetPost(id string) (*entity.Post, error) {
	var post entity.Post
	if err := r.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) GetPosts() (*[]entity.Post, error) {
	var posts []entity.Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return &posts, nil
}
