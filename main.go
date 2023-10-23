package main

import (
	"log"
	"net/http"
	"todogorest/config"
	"todogorest/controllers"
	"todogorest/data/response"
	"todogorest/helpers"
	"todogorest/models"
	"todogorest/repositories"
	"todogorest/routes"
	"todogorest/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := config.ConnectToDB()

	db.Table("todos").AutoMigrate(&models.Todo{})

	db.Table("users").AutoMigrate(&models.User{})

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{StatusCode: http.StatusNotFound, Code: helpers.NotFound, Data: response.ErrorMessage{Message: "Route not found"}})
	})

	// Todo repositories
	todoRepositories := repositories.NewTodoRepository(db)
	userRepositories := repositories.NewUserRepository(db)

	// Services
	todoServices := services.NewTodoServicesImpl(todoRepositories)
	userServices := services.NewUserServicesImpl(userRepositories)

	// Controllers
	todoController := controllers.NewTodoController(todoServices)
	userController := controllers.NewUserController(userServices)

	// Routes
	routes := routes.NewRouter(todoController, userController)

	server := &http.Server{
		Addr:    ":5051",
		Handler: routes,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err.Error())
	}
}
