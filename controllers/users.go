package controllers

import (
	"github.com/akhidnukhlis/entity"
	"github.com/akhidnukhlis/models"
	"github.com/labstack/echo"
	"net/http"
)

func StoreUser(c echo.Context) error {
	user := new(entity.UserRequest)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
	}

	result, err := models.StoreUser(user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
