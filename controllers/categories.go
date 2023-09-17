package controllers

import (
	"github.com/akhidnukhlis/common"
	"github.com/akhidnukhlis/entity"
	"github.com/akhidnukhlis/models"
	"github.com/labstack/echo"
	"net/http"
)

func FetchAllCategory(c echo.Context) error {
	result, err := models.FetchAllCategory()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetByIdCategory(c echo.Context) error {
	id := c.Param("id")

	valid := common.IsValidUUID(id)
	if valid == false {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": "UUID no valid"})
	}

	result, err := models.GetByIdCategory(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreCategory(c echo.Context) error {
	category := new(entity.CategoryRequest)
	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
	}

	result, err := models.StoreCategory(category)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateCategory(c echo.Context) error {
	id := c.Param("id")

	valid := common.IsValidUUID(id)
	if valid == false {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": "UUID no valid"})
	}

	category := new(entity.CategoryRequest)
	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
	}

	result, err := models.UpdateCategory(id, category)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteCategory(c echo.Context) error {
	id := c.Param("id")

	valid := common.IsValidUUID(id)
	if valid == false {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": "UUID no valid"})
	}

	_, err := models.GetByIdCategory(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
	}

	result, err := models.DeleteCategory(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
