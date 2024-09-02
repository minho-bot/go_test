package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	postEntity "go_test/src/domain/post/entity"
	userEntity "go_test/src/domain/user/entity"
)

// SetupDB는 데이터베이스를 설정하고 연결을 반환합니다.
func SetupDB() *gorm.DB {
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

	// 데이터베이스 마이그레이션 (선택적)
	if err := db.AutoMigrate(&userEntity.User{}, &postEntity.Post{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
