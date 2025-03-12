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
	deleteUserUseCase := usecases.NewDeleteUserUseCase(userRepository)
	getUserEmailUseCase := usecases.NewGetUserEmailUseCase(userRepository)
	updateUserUseCase := usecases.NewUpdateUserUseCase(userRepository)

	userHandler := handlers.NewUserHandler(createUserUseCase, deleteUserUseCase, updateUserUseCase, getUserEmailUseCase)

	r := gin.Default()
	r.POST("/user", userHandler.CreateUser)
	r.DELETE("/user/:id", userHandler.DeleteUser)
	r.PUT("/user/:id", userHandler.UpdateUser)
	r.GET("/user/:email", userHandler.GetUserEmail)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
