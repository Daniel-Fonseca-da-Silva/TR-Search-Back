package main

import (
	"log"
	"os"

	"tr-search-back/internal/domain/user"
	"tr-search-back/internal/handlers"
	"tr-search-back/internal/usecases"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := handlers.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&user.User{})
	userRepository := handlers.NewUserRepository(db)
	createUserUseCase := usecases.NewUseCase(userRepository)
	userHandler := handlers.NewUserHandler(createUserUseCase)

	r := gin.Default()
	r.POST("/users", userHandler.CreateUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
