package routes

import (
	"todogorest/controllers"
	"todogorest/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter(todoController *controllers.TodoController, userController *controllers.UserController, verificationCodeController *controllers.VerificationCodeController) *gin.Engine {
	router := gin.Default()

	baseRouter := router.Group("/api")

	//AUTH
	authRouter := baseRouter.Group("/auth")

	authRouter.POST("/signin", userController.Signin)

	authRouter.POST("/signup", userController.Signup)

	authRouter.POST("/refresh", userController.RefreshUser)

	authRouter.PUT("/verify/:user_id", userController.VerifyUser)

	authRouter.PUT("/verify/:user_id/resend", userController.ResendVerification)

	// VERIFICATION CODES ROUTES
	verificationCodeRouter := baseRouter.Group("/verification-codes")

	verificationCodeRouter.Use(middlewares.IsAuth)

	verificationCodeRouter.POST("/:user_id", verificationCodeController.CreateVerificationCode)

	verificationCodeRouter.GET("/:id", verificationCodeController.GetVerificationCode)

	verificationCodeRouter.PUT("/:user_id", verificationCodeController.UpdateVerificationCode)

	verificationCodeRouter.DELETE("/:id", verificationCodeController.DeleteVerificationCode)

	// LOGGED IN USER
	loggedInAuthRoutes := authRouter.Group("/me")

	loggedInAuthRoutes.Use(middlewares.IsAuth)

	loggedInAuthRoutes.GET("/", userController.GetSignedInUser)

	loggedInAuthRoutes.PUT("/")

	//USERS
	usersRoutes := baseRouter.Group("/users")

	usersRoutes.Use(middlewares.IsAuth)

	usersRoutes.GET("/:user_id", userController.GetUserById)

	// TODOS
	todoRouter := baseRouter.Group("/todos")

	todoRouter.Use(middlewares.IsAuth)

	todoRouter.GET("/", todoController.GetAllTodos)

	todoRouter.POST("/", todoController.CreateTodo)

	todoRouter.DELETE("/:todo_id", todoController.DeleteTodo)

	todoRouter.GET("/:todo_id", todoController.GetTodo)

	todoRouter.PUT("/:todo_id", todoController.UpdateTodo)

	return router
}
