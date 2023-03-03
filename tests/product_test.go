package tests

import (
	_ "testing"
)

type Product struct {
	Id   		int `validate:"required"`
	Name  		string `validate:"required"`
	Price 		int `validate:"required"`
}
