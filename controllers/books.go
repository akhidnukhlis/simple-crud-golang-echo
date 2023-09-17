package controllers

import (
	"github.com/akhidnukhlis/common"
	"github.com/akhidnukhlis/entity"
	"github.com/akhidnukhlis/models"
	"github.com/labstack/echo"
	"net/http"
)

func FetchAllBook(c echo.Context) error {
	result, err := models.FetchAllBook()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetByIdBook(c echo.Context) error {
	id := c.Param("id")

	valid := common.IsValidUUID(id)
	if valid == false {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": "UUID no valid"})
	}

	result, err := models.GetByIdBook(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreBook(c echo.Context) error {
	book := new(entity.BookRequest)
	if err := c.Bind(book); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
	}

	result, err := models.StoreBook(book)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateBook(c echo.Context) error {
	id := c.Param("id")

	valid := common.IsValidUUID(id)
	if valid == false {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": "UUID no valid"})
	}

	_, err := models.GetByIdBook(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
	}

	book := new(entity.BookRequest)
	if err := c.Bind(book); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
	}

	result, err := models.UpdateBook(id, book)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteBook(c echo.Context) error {
	id := c.Param("id")

	valid := common.IsValidUUID(id)
	if valid == false {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": "UUID no valid"})
	}

	_, err := models.GetByIdBook(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
	}

	result, err := models.DeleteBook(id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
