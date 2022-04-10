package routes

import (
	"MyMiniProject/constant"
	c "MyMiniProject/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.POST("/auth/login", c.LoginController)
	e.POST("/users", c.CreateUserController)

	// Route Books
	e.GET("/motivation", c.GetAllMotivationController)
	e.GET("/motivation/:id", c.GetMotivationOneController)

	// JWT
	jwtAuth := e.Group("/jwt/redirected")
	jwtAuth.Use(middleware.JWT([]byte(constant.SECRET_JWT)))

	// Route with JWT Auth
	jwtAuth.GET("/users", c.GetAllUsersController)
	jwtAuth.GET("/users/:id", c.GetUserController)
	jwtAuth.DELETE("/users/:id", c.DeleteUserController)
	jwtAuth.PUT("/users/:id", c.UpdateUserController)

	jwtAuth.POST("/books", c.CreateMotivationController)
	jwtAuth.DELETE("/books/:id", c.DeleteMotivationController)
	jwtAuth.PUT("/books/:id", c.UpdateMotivationController)
	return e
}
