package main

import (
	"devlog/config"
	"devlog/internal/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
)

func main() {

	// 환경 설정 로드
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// DB 인스턴스 얻기
	db, err := config.GetDBInstance(conf)
	if err != nil {
		log.Fatal("Error getting DB instance:", err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	//// DB 마이그레이션 (User 모델 자동 생성)
	//if err := db.AutoMigrate(&models.User{}); err != nil {
	//	log.Fatal("Error migrating database:", err)
	//}

	// Gin 라우터 설정
	r := gin.Default()

	// DB 인스턴스를 핸들러에 전달
	userHandler := handlers.NewUserHandler(db)

	// 라우팅 설정
	r.POST("/user", userHandler.CreateUser)
	r.GET("/hello", func(c *gin.Context) {
		fmt.Fprintln(c.Writer, "Hello World")
	})
	// 서버 실행
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start the server:", err)
	}

	fmt.Println("Server is running on http://localhost:8080")
}
