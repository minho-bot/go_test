package resolver

import (
	"context"
	"go_test/graph/gql_model"
	"go_test/src/domain/user/service"
)

type UserResolver struct {
	userService *service.UserService
}

func NewUserResolver(service *service.UserService) *UserResolver {
	return &UserResolver{userService: service}
}

func (r *UserResolver) CreateUser(ctx context.Context, name string, email string) (*gql_model.User, error) {
	user, err := r.userService.CreateUser(ctx, name, email)
	if err != nil {
		return nil, err
	}
	return user.ToGraphQLModel(), nil
}

func (r *UserResolver) UpdateUser(ctx context.Context, id string, name *string, email *string) (*gql_model.User, error) {
	user, err := r.userService.UpdateUser(ctx, id, name, email)
	if err != nil {
		return nil, err
	}
	return user.ToGraphQLModel(), nil
}

func (r *UserResolver) DeleteUser(ctx context.Context, id string) (*gql_model.User, error) {
	user, err := r.userService.DeleteUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return user.ToGraphQLModel(), nil
}

func (r *UserResolver) User(ctx context.Context, id string) (*gql_model.User, error) {
	user, err := r.userService.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return user.ToGraphQLModel(), nil
}

func (r *UserResolver) Users(ctx context.Context) ([]*gql_model.User, error) {
	users, err := r.userService.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	var gqlUsers []*gql_model.User
	for _, user := range *users {
		gqlUsers = append(gqlUsers, user.ToGraphQLModel())
	}

	return gqlUsers, nil
}
