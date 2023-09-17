package common_test

import (
	"github.com/akhidnukhlis/common"
	"testing"
)

func TestNewQueryBuilder_Where(t *testing.T) {
	query := common.NewQueryBuilder("users").Select("*").Where("id = 13056b87-2fc9-48e3-b021-deea24dee9c3").Where("username = 'Rafaela'").Limit(1).Build()
	expected := "SELECT * FROM users WHERE id = '13056b87-2fc9-48e3-b021-deea24dee9c3' AND username = 'Rafaela' LIMIT 1"
	if query != expected {
		t.Errorf("expected %s, got %s", expected, query)
	}
}

func TestNewQueryBuilder_EntityWhereStruct(t *testing.T) {
	type User struct {
		ID       string `json:"id,omitempty"`
		Username string `json:"username,omitempty"`
		Age      int    `json:"age,omitempty"`
	}

	var entity User

	structParams := struct {
		ID        string `json:"id,omitempty"`
		Username  string `json:"username,omitempty"`
		Age       int    `json:"age,omitempty"`
		StartDate string `json:"start_date,omitempty"`
		EndDate   string `json:"end_date,omitempty"`
		Q         string `json:"q,omitempty"`
	}{
		ID:        "13056b87-2fc9-48e3-b021-deea24dee9c3",
		Username:  "Rafaela",
		StartDate: "2020-01-01",
		EndDate:   "2020-01-31",
		Q:         "test",
	}

	query := common.NewQueryBuilder("users").Entity(entity).Select("*").WhereStruct(structParams).Build()
	expected := "SELECT * FROM users WHERE username = 'Rafaela' AND created_at >= '2020-01-01' AND created_at <= '2020-01-31' AND id = '13056b87-2fc9-48e3-b021-deea24dee9c3'"
	if query != expected {
		t.Errorf("expected %s, got %s", expected, query)
	}
}

func TestNewQueryBuilder_WhereStruct(t *testing.T) {
	structParams := struct {
		ID        string `json:"id,omitempty"`
		Username  string `json:"username,omitempty"`
		Age       int    `json:"age,omitempty"`
		StartDate string `json:"start_date,omitempty"`
		EndDate   string `json:"end_date,omitempty"`
	}{
		ID:        "13056b87-2fc9-48e3-b021-deea24dee9c3",
		Username:  "Rafaela",
		StartDate: "2020-01-01",
		EndDate:   "2020-01-31",
	}

	query := common.NewQueryBuilder("users").Select("*").WhereStruct(structParams).Build()
	expected := "SELECT * FROM users WHERE username = 'Rafaela' AND created_at >= '2020-01-01' AND created_at <= '2020-01-31' AND id = '13056b87-2fc9-48e3-b021-deea24dee9c3'"
	if query != expected {
		t.Errorf("expected %s, got %s", expected, query)
	}
}

func TestNewQueryBuilder_WhereStruct_Recursive(t *testing.T) {
	type User struct {
		ID       string `json:"id,omitempty"`
		Username string `json:"username,omitempty"`
		Age      int    `json:"age,omitempty"`
		Foo      int    `json:"foo,omitempty"`
		Bar      string `json:"bar,omitempty"`
	}

	structData := User{
		ID:       "13056b87-2fc9-48e3-b021-deea24dee9c3",
		Username: "Rafaela",
		Age:      20,
	}

	structMeta := struct {
		StartDate string `json:"start_date,omitempty"`
		EndDate   string `json:"end_date,omitempty"`
	}{
		StartDate: "2020-01-01",
		EndDate:   "2020-01-31",
	}

	structParams := struct {
		Data interface{} `json:"data,omitempty"`
		Meta interface{} `json:"meta,omitempty"`
		Foo  int         `json:"foo,omitempty"`
		Bar  string      `json:"bar,omitempty"`
	}{
		Data: structData,
		Meta: structMeta,
		Foo:  1,
		Bar:  "test",
	}

	query := common.NewQueryBuilder("users").Entity(User{}).Select("*").WhereStruct(structParams).Build()
	expected := "SELECT * FROM users WHERE id = '13056b87-2fc9-48e3-b021-deea24dee9c3' AND username = 'Rafaela' AND created_at >= '2020-01-01' AND created_at <= '2020-01-31' AND foo = 1 AND bar = 'test'"
	if query != expected {
		t.Errorf("expected %s, got %s", expected, query)
	}
}

func TestNewQueryBuilder_Insert(t *testing.T) {
	query := common.NewQueryBuilder("transactions").Insert(struct {
		ID       string `json:"id,omitempty"`
		Username string `json:"username,omitempty"`
	}{
		ID:       "13056b87-2fc9-48e3-b021-deea24dee9c3",
		Username: "Rafaela",
	}).Build()
	expected := "INSERT INTO users(id, username) VALUES ('13056b87-2fc9-48e3-b021-deea24dee9c3', 'Rafaela')"
	if query != expected {
		t.Errorf("expected %s, got %s", expected, query)
	}
}

func TestNewQueryBuilder_Update(t *testing.T) {
	query := common.NewQueryBuilder("transactions").Update(struct {
		ID       string `json:"id,omitempty"`
		Username string `json:"username,omitempty"`
	}{
		ID:       "13056b87-2fc9-48e3-b021-deea24dee9c3",
		Username: "Rafaela",
	}).Where("id = '13056b87-2fc9-48e3-b021-deea24dee9c3'").Build()
	expected := "UPDATE users SET id = '13056b87-2fc9-48e3-b021-deea24dee9c4', username = 'Rafaela' WHERE id = '13056b87-2fc9-48e3-b021-deea24dee9c3'"
	if query != expected {
		t.Errorf("expected %s, got %s", expected, query)
	}
}

func TestNewQueryBuilder_Delete(t *testing.T) {
	query := common.NewQueryBuilder("users").Delete().Where("id = '13056b87-2fc9-48e3-b021-deea24dee9c3'").Build()
	expected := "DELETE FROM users WHERE id = '13056b87-2fc9-48e3-b021-deea24dee9c3'"
	if query != expected {
		t.Errorf("expected %s, got %s", expected, query)
	}
}

func TestNewQueryBuilder_In(t *testing.T) {
	type User struct {
		ID       string `json:"id,omitempty"`
		Username string `json:"username,omitempty"`
	}

	structData := User{
		ID:       "13056b87-2fc9-48e3-b021-deea24dee9c3",
		Username: "Rafaela",
	}

	structMeta := struct {
		Groups []string `json:"groups,omitempty"`
	}{
		Groups: []string{"A", "B", "C"},
	}

	structParams := struct {
		Data interface{} `json:"data,omitempty"`
		Meta interface{} `json:"meta,omitempty"`
	}{
		Data: structData,
		Meta: structMeta,
	}

	query := common.NewQueryBuilder("users").Entity(User{}).Select("*").WhereStruct(structParams).In("groups", structMeta.Groups).Build()
	expected := "SELECT * FROM users WHERE id = '13056b87-2fc9-48e3-b021-deea24dee9c3' AND username = 'Rafaela' AND groups IN ('A', 'B', 'C')"
	if query != expected {
		t.Errorf("expected %s, got %s", expected, query)
	}
}
