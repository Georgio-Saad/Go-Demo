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

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{StatusCode: http.StatusNotFound, Code: helpers.NotFound, Data: response.ErrorMessage{Message: "Route not found"}})
	})

	// Todo repositories
	todoRepositories := repositories.NewTodoRepository(db)

	// Services
	todoServices := services.NewTodoServicesImpl(todoRepositories)

	// Controllers
	todoController := controllers.NewTodoController(todoServices)

	// Routes
	routes := routes.NewRouter(todoController)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Okay"})
	})

	server := &http.Server{
		Addr:    ":5051",
		Handler: routes,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err.Error())
	}
}
