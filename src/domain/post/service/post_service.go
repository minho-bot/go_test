package service

import (
	"go_test/src/domain/post/entity"
	"go_test/src/domain/post/repository"
)

type PostService struct {
	postRepository *repository.PostRepository
}

func NewPostService(repository *repository.PostRepository) *PostService {
	return &PostService{postRepository: repository}
}

func (s *PostService) CreatePost(title, content, author string) (*entity.Post, error) {
	newPost := entity.Post{
		Title:   title,
		Author:  author,
		Content: content,
	}

	return s.postRepository.CreatePost(&newPost)
}

func (s *PostService) UpdatePost(id string, title, content *string) (*entity.Post, error) {
	post, err := s.postRepository.GetPost(id)
	if err != nil {
		return nil, err
	}

	if title != nil {
		post.Title = *title
	}
	if content != nil {
		post.Content = *content
	}

	return s.postRepository.UpdatePost(post)
}

func (s *PostService) DeletePost(id string) (*entity.Post, error) {
	post, err := s.postRepository.GetPost(id)
	if err != nil {
		return nil, err
	}

	if err := s.postRepository.DeletePost(post); err != nil {
		return nil, err
	}

	return post, nil
}

func (s *PostService) GetPost(id string) (*entity.Post, error) {
	return s.postRepository.GetPost(id)
}

func (s *PostService) GetPosts() (*[]entity.Post, error) {
	return s.postRepository.GetPosts()
}
