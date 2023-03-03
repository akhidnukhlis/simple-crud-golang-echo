package models

import (
	"github.com/akhidnukhlis/db"
	validator "github.com/go-playground/validator"
	"net/http"
)

type Product struct {
	Id      	int    	`json:"id"`
	Name    	string 	`json:"name" validate:"required"`
	Price    	int		`json:"price" validate:"required"`
}

func FetchAllProduct() (Response, error) {
	var obj Product
	var arrobj []Product
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM product"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Price)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func StoreProduct(name string, price int) (Response, error) {
	var res Response

	v := validator.New()

	prod := Product{
		Name		: name,
		Price		: price,
	}

	err := v.Struct(prod)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT product (name, price) VALUES (?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, price)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdateProduct(id int, name string, price int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE product SET name = ?, price = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, price, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteProduct(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM product WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
