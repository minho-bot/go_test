package config

import (
	"go_test/graph/resolvers"

	postRepository "go_test/src/domain/post/repository"
	postResolver "go_test/src/domain/post/resolver"
	postService "go_test/src/domain/post/service"

	userRepository "go_test/src/domain/user/repository"
	userResolver "go_test/src/domain/user/resolver"
	userService "go_test/src/domain/user/service"

	"gorm.io/gorm"
)

// SetupDependencies는 DB를 설정하고 의존성을 초기화합니다.
func SetupDependencies(db *gorm.DB) *resolvers.Resolver {

	postRepository := postRepository.NewPostRepository(db)
	postService := postService.NewPostService(postRepository)
	postResolver := postResolver.NewPostResolver(postService)

	userRepository := userRepository.NewUserRepository(db)
	userService := userService.NewUserService(userRepository)
	userResolver := userResolver.NewUserResolver(userService)

	return resolvers.NewResolver(
		postResolver,
		userResolver,
	)
}
