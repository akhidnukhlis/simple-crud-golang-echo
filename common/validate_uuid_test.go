package common_test

import (
	"github.com/akhidnukhlis/common"
	"testing"
)

func TestIsValidUUID(t *testing.T) {
	uuid := "4c931f11-0085-4586-b52c-67516eb5cc19"
	snake := common.IsValidUUID(uuid)
	expected := true
	if snake != expected {
		t.Errorf("expected %v, got %v", expected, snake)
	}
}
