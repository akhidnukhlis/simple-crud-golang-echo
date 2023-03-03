package controllers

import (
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type Customer struct {
	Name   		string `validate:"required"`
	Email  		string `validate:"required,email"`
	Address 	string `validate:"required"`
	Age   		int    `validate:"gte=17,lte=35"`
}

func TestVariableValidation(c echo.Context) error {
	v := validator.New()

	email := "nukhlis@gmail.com"

	err := v.Var(email, "required,email")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Email not valid",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Success"})
}

func TestStructValidation(c echo.Context) error {
	v := validator.New()

	cust := Customer{
		Name	:   "Akhid Nukhlis",
		Email	:  	"nukhlis@gmail.com",
		Address	: 	"jakarta",
		Age		:   25,
	}

	err := v.Struct(cust)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Success"})
}
