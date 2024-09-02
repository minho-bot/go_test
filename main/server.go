package main

import (
	"go_test/graph/generated"
	"log"
	"os"

	"go_test/src/config"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {

	// DB 초기화
	db := config.SetupDB()

	// 의존성 주입 및 데이터베이스 초기화
	resolver := config.SetupDependencies(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Gin 라우터 인스턴스 생성
	r := gin.Default()

	// GraphQL 서버 핸들러 생성
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver, // Resolver를 설정
	}))

	// Playground 핸들러 설정
	r.GET("/", func(c *gin.Context) {
		playground.Handler("GraphQL playground", "/query")(c.Writer, c.Request)
	})

	// GraphQL 엔드포인트 설정
	r.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	// 서버 실행
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	r.Run(":" + port) // Gin의 Run 메소드는 내부적으로 log.Fatal을 사용하여 서버를 시작합니다.
}
