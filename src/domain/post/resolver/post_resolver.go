package resolver

import (
	"context"
	"go_test/graph/gql_model"
	"go_test/src/domain/post/service"
)

type PostResolver struct {
	postService *service.PostService
}

func NewPostResolver(service *service.PostService) *PostResolver {
	return &PostResolver{postService: service}
}

func (r *PostResolver) CreatePost(ctx context.Context, title string, content string, author string) (*gql_model.Post, error) {
	post, err := r.postService.CreatePost(title, content, author)
	if err != nil {
		return nil, err
	}
	return post.ToGraphQLModel(), nil
}

func (r *PostResolver) UpdatePost(ctx context.Context, id string, title *string, content *string) (*gql_model.Post, error) {
	post, err := r.postService.UpdatePost(id, title, content)
	if err != nil {
		return nil, err
	}
	return post.ToGraphQLModel(), nil
}

func (r *PostResolver) DeletePost(ctx context.Context, id string) (*gql_model.Post, error) {
	post, err := r.postService.DeletePost(id)
	if err != nil {
		return nil, err
	}
	return post.ToGraphQLModel(), nil
}

func (r *PostResolver) Post(ctx context.Context, id string) (*gql_model.Post, error) {
	post, err := r.postService.GetPost(id)
	if err != nil {
		return nil, err
	}
	return post.ToGraphQLModel(), nil
}

func (r *PostResolver) Posts(ctx context.Context) ([]*gql_model.Post, error) {
	posts, err := r.postService.GetPosts()
	if err != nil {
		return nil, err
	}

	var gqlPosts []*gql_model.Post
	for _, post := range *posts {
		gqlPosts = append(gqlPosts, post.ToGraphQLModel())
	}

	return gqlPosts, nil
}
