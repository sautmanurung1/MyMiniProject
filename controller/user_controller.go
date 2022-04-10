package controller

import (
	"MyMiniProject/config"
	"MyMiniProject/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAllUsersController(c echo.Context) error {
	DB := config.Connect()
	var users []models.User

	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	DB := config.Connect()
	id := c.Param("id")
	user := models.User{}

	if err := DB.Where("ID = ?", id).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "Users Is Empty in Database",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success Get User",
		"User":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	DB := config.Connect()
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		return err
	}

	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	DB := config.Connect()
	id := c.Param("id")

	DB.Delete(&models.User{}, id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success Delete Users",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	DB := config.Connect()
	id := c.Param("id")
	user := models.User{}

	DB.Where("ID = ?", id).First(&user)

	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "Users Is Empty in Database",
		})
	}

	payload := models.User{}
	err := c.Bind(&payload)
	if err != nil {
		return err
	}

	user.Name = payload.Name
	user.Email = payload.Email
	user.Password = payload.Password
	DB.Save(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success Update User",
		"Books":   user,
	})
}
