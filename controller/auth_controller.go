package controller

import (
	"MyMiniProject/config"
	"MyMiniProject/middleware"
	"MyMiniProject/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func LoginController(c echo.Context) error {
	DB := config.Connect()

	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		return err
	}

	err2 := DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Email or Password is Wrong",
		})
	}

	token, e := middleware.CreateToken(strconv.Itoa(int(user.ID)), user.Name)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Error Create Token",
		})
	}

	responseData := models.Auth{}
	responseData.Name = user.Name
	responseData.Email = user.Email
	responseData.Token = token

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":       "Success Login",
		"Response Data": responseData,
	})
}
