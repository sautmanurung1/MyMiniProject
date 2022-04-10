package controller

import (
	"MyMiniProject/config"
	"MyMiniProject/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateMotivationController(c echo.Context) error {
	DB := config.Connect()
	motivation := models.Motivation{}

	err := c.Bind(&motivation)
	if err != nil {
		return err
	}

	if err := DB.Save(&motivation).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success Create New Motivation",
	})
}

func GetAllMotivationController(c echo.Context) error {
	DB := config.Connect()
	var motivation []models.Motivation

	if err := DB.Find(&motivation).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":    "Success Get All Motivation",
		"Motivation": motivation,
	})
}

func GetMotivationOneController(c echo.Context) error {
	DB := config.Connect()
	ID := c.Param("id")
	motivation := models.Motivation{}

	DB.Where("ID = ?", ID).First(&motivation)

	if motivation.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "Motivation Is Empty in Database",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":    "Success Get Motivation",
		"Motivation": motivation,
	})
}

func UpdateMotivationController(c echo.Context) error {
	DB := config.Connect()
	ID := c.Param("id")
	motivation := models.Motivation{}

	DB.Where("ID = ?", ID).First(&motivation)

	if motivation.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"Message": "Books Not Found In Database",
		})
	}

	payload := models.Motivation{}
	err := c.Bind(&payload)
	if err != nil {
		return err
	}

	motivation.ID = payload.ID
	motivation.Motivation = payload.Motivation
	DB.Save(&motivation)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success Update Books",
		"Books":   motivation,
	})
}

func DeleteMotivationController(c echo.Context) error {
	DB := config.Connect()
	id := c.Param("id")

	DB.Delete(&models.Motivation{}, id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success Delete Books",
	})
}
