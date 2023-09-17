package tests

import (
	"fmt"
	"testing"

	validator "github.com/go-playground/validator/v10"
)

type Customer struct {
	Name    string `validate:"required"`
	Email   string `validate:"required,email"`
	Address string `validate:"required"`
	Age     int    `validate:"gte=17,lte=35"`
}

func TestVariableValidation(t *testing.T) {
	v := validator.New()

	email := "nukhlis@gmail.com"

	err := v.Var(email, "required,email")
	if err != nil {
		fmt.Println("validation email got", err)
	}

	fmt.Println("success validation email")
}

func TestStructValidation(t *testing.T) {
	v := validator.New()

	cust := Customer{
		Name:    "Akhid Nukhlis",
		Email:   "nukhlis@gmail.com",
		Address: "jakarta",
		Age:     25,
	}
	err := v.Struct(cust)
	if err != nil {
		fmt.Println("validation struct got:", err)
	}

	fmt.Println("success validation struct")
}
