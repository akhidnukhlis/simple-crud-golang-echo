package models

import (
	"fmt"
	"github.com/akhidnukhlis/entity"
	"github.com/akhidnukhlis/helpers"
	"github.com/jinzhu/gorm"

	"github.com/akhidnukhlis/db"
)

func CheckLogin(username, password string) (bool, error) {
	con := db.CreateCon()

	var user entity.User
	result := con.Where("username = $1", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// User with the given username not found
			fmt.Println("Username not found")
			return false, nil
		}
		// Some other error occurred
		return false, result.Error
	}

	match, err := helpers.CheckPasswordHash(password, user.Password)
	if !match {
		fmt.Println("Hash and password doesn't match.")
		return false, err
	}

	return true, nil
}
