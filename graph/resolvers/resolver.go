package resolvers

import (
	post "go_test/src/domain/post/resolver"
	user "go_test/src/domain/user/resolver"
)

type Resolver struct {
	postResolver *post.PostResolver
	userResolver *user.UserResolver
}

func NewResolver(postResolver *post.PostResolver, userResolver *user.UserResolver) *Resolver {
	return &Resolver{postResolver: postResolver, userResolver: userResolver}
}
