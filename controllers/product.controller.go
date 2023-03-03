package controllers

import (
	"github.com/akhidnukhlis/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func FetchAllProduct(c echo.Context) error {
	result, err := models.FetchAllProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreProduct(c echo.Context) error {
	name 		:= c.FormValue("name")
	price 		:= c.FormValue("price")

	convPrice, err := strconv.Atoi(price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreProduct(name, convPrice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateProduct(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	price := c.FormValue("price")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	conv_price, err := strconv.Atoi(price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateProduct(conv_id, name, conv_price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteProduct(c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteProduct(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
