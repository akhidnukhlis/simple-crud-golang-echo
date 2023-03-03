package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/akhidnukhlis/helpers"

	"github.com/akhidnukhlis/db"
)

type Users struct {
	UserId    		int    `json:"userID"`
	Username 		string `json:"usename"`
	Password 		string `json:"password"`
	FullName 		string `json:"fullName"`
	Email 			string `json:"email"`
	RoleCode 		string `json:"roleCode"`
	CreatedDate 	time.Time `json:"createdDate"`
	ModifiedDate 	time.Time `json:"modifiedDate"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj Users
	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.UserId, &obj.Username, &pwd, &obj.FullName, &obj.Email, &obj.RoleCode, &obj.CreatedDate, &obj.ModifiedDate,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and password doesn't match.")
		return false, err
	}

	return true, nil
}
