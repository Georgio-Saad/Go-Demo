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
	"github.com/joho/godotenv"
)

func init() {

}

func main() {
	r := gin.Default()

	dotEnvErr := godotenv.Load(".env")

	if dotEnvErr != nil {
		panic(dotEnvErr)
	}

	db := config.ConnectToDB()

	db.Table("products").AutoMigrate(&models.Product{})

	db.Table("todos").AutoMigrate(&models.Todo{})

	db.Table("users").AutoMigrate(&models.User{})

	db.Table("verification_codes").AutoMigrate(&models.VerificationCode{})

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{StatusCode: http.StatusNotFound, Code: helpers.NotFound, Data: response.ErrorMessage{Message: "Route not found"}})
	})

	// Todo repositories
	todoRepositories := repositories.NewTodoRepository(db)
	userRepositories := repositories.NewUserRepository(db)
	verificationCodeRepositories := repositories.NewVerificationCodeRepository(db)
	productRepositories := repositories.NewProductRepositoryImpl(db)

	// Services
	todoServices := services.NewTodoServicesImpl(todoRepositories)
	userServices := services.NewUserServicesImpl(userRepositories, verificationCodeRepositories)
	verificationCodeServices := services.NewVerificationCodeServicesImpl(verificationCodeRepositories)
	productServices := services.NewProductServicesImpl(productRepositories)

	// Controllers
	generalController := controllers.NewGeneralController()
	todoController := controllers.NewTodoController(todoServices)
	userController := controllers.NewUserController(userServices)
	verificationCodeController := controllers.NewVerificationCodeController(verificationCodeServices)
	productController := controllers.NewProductController(productServices)

	// Routes
	routes := routes.NewRouter(todoController, userController, verificationCodeController, generalController, productController)

	server := &http.Server{
		Addr:    ":5051",
		Handler: routes,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err.Error())
	}
}
