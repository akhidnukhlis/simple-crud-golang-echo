package models

import (
	"fmt"
	"github.com/akhidnukhlis/db"
	"github.com/akhidnukhlis/entity"
	"github.com/akhidnukhlis/helpers"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func StoreUser(payload *entity.UserRequest) (*Response, error) {
	var (
		res Response
		con = db.CreateCon()
	)

	hash, _ := helpers.HashPassword(payload.Password)

	// Create a new User instance with the payload data
	user := &entity.User{
		ID:        uuid.New().String(),
		Username:  payload.Username,
		Password:  hash,
		FullName:  payload.FullName,
		Email:     payload.Email,
		RoleCode:  payload.RoleCode,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Insert the new User record into the database
	if err := con.Create(user).Error; err != nil {
		err = fmt.Errorf("creating user got: %w", err)

		return nil, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = user

	return &res, nil
}
