package routes

import (
	"todogorest/controllers"
	"todogorest/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter(todoController *controllers.TodoController, userController *controllers.UserController) *gin.Engine {
	router := gin.Default()

	baseRouter := router.Group("/api")

	//AUTH
	authRouter := baseRouter.Group("/auth")

	authRouter.POST("/signin", userController.Signin)

	authRouter.POST("/signup", userController.Signup)

	authRouter.POST("/refresh", userController.RefreshUser)

	// VERIFICATION CODES ROUTES
	// verificationCodeRouter := baseRouter.Group("/verification")

	// verificationCodeRouter.POST("/:user_id")

	// verificationCodeRouter.GET("/:user_id")

	// verificationCodeRouter.GET("/:verification_code_id")

	// verificationCodeRouter.PUT("/:user_id")

	// verificationCodeRouter.DELETE("/:user_id")

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
