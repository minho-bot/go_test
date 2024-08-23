package main

import (
	"go_test/graph/generated"
	"go_test/graph/resolvers"
	postEntity "go_test/src/domain/post/entity"
	userEntity "go_test/src/domain/user/entity"
	"log"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const defaultPort = "8080"

func main() {
	// 데이터베이스 연결 DSN (Data Source Name)
	dsn := "user:password@tcp(localhost:3306)/userdb?charset=utf8mb4&parseTime=True&loc=Local"

	// Logger 설정
	logger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
		ParameterizedQueries:      true,
	})

	// 데이터베이스 연결
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 데이터베이스 연결 상태 확인 (선택적)
	sqlDB, dbErr := db.DB()
	if dbErr != nil {
		log.Fatalf("Failed to get database handle: %v", dbErr)
	}
	defer sqlDB.Close() // 데이터베이스 연결 종료 보장

	// 데이터베이스 마이그레이션
	if err := db.AutoMigrate(&userEntity.User{}, &postEntity.Post{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Gin 라우터 인스턴스 생성
	r := gin.Default()

	// GraphQL 서버 핸들러 생성
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolvers.Resolver{
			DB: db, // 데이터베이스 연결을 Resolver 구조체에 전달
		},
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
