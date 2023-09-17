package common_test

import (
	"github.com/akhidnukhlis/common"
	"testing"
)

func TestStringPascalToSnake(t *testing.T) {
	snake := common.StringPascalToSnake("TransactionRepository")
	expected := "transaction_repository"
	if snake != expected {
		t.Errorf("expected %s, got %s", expected, snake)
	}
}

func TestStringInSlice_True(t *testing.T) {
	slice := []string{"a", "b", "c"}
	searchString := "a"
	expected := true
	result := common.StringInSlice(slice, searchString)
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestStringInSlice_False(t *testing.T) {
	slice := []string{"a", "b", "c"}
	searchString := "d"
	expected := false
	result := common.StringInSlice(slice, searchString)
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
