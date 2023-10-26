package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
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
	"github.com/kataras/jwt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"golang.org/x/text/language"
)

//go:generate todoappprest extract -sourceLanguage en

func main() {
	r := gin.Default()

	dotEnvErr := godotenv.Load(".env")

	if dotEnvErr != nil {
		panic(dotEnvErr)
	}

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	localizer := i18n.NewLocalizer(bundle, "en")

	buying := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "BuyingCookies",
			One:   "Youre buying 1 cookie",
			Other: "Youre buying {{.PluralCount}} cookies",
		},
		PluralCount: 5,
	})

	fmt.Printf("%s\n", buying)

	blocklist := jwt.NewBlocklist(2 * 7 * 24 * time.Hour)

	db := config.ConnectToDB()

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Todo{}, &models.VerificationCode{})

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
	userServices := services.NewUserServicesImpl(userRepositories, verificationCodeRepositories, blocklist)
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
