package models

import (
	"fmt"
	"github.com/akhidnukhlis/db"
	"github.com/akhidnukhlis/entity"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

func FetchAllCategory() (Response, error) {
	var (
		c   entity.Categories
		res Response

		con = db.CreateCon()
	)

	sqlStatement := "SELECT * FROM categories WHERE status = true"

	con.Raw(sqlStatement).Scan(&c)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = c

	return res, nil
}

func GetByIdCategory(id string) (*Response, error) {
	var (
		category entity.Category
		res      Response

		con = db.CreateCon()
	)

	// Assuming "con" is your GORM database connection, and "id" is the ID you want to query.
	err := con.Where("id = $1 AND status = true", id).First(&category).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Handle the case when the Category is not found
			err = fmt.Errorf("getting book by id got: %w", err)

			return nil, err
		} else {
			// Handle other query errors
			err = fmt.Errorf("Query error: %w", err)

			return nil, err
		}
		return nil, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = category

	return &res, nil
}

func StoreCategory(payload *entity.CategoryRequest) (*Response, error) {
	var (
		res Response
		con = db.CreateCon()
	)

	// Create a new Category instance with the payload data
	category := &entity.Category{
		ID:        uuid.New().String(),
		Name:      payload.Name,
		Status:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Insert the new Category record into the database
	if err := con.Create(category).Error; err != nil {
		err = fmt.Errorf("creating category got: %w", err)

		return nil, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = category

	return &res, nil
}

func UpdateCategory(id string, payload *entity.CategoryRequest) (Response, error) {
	var (
		res Response
		con = db.CreateCon()
	)

	// Check if the Category with the given ID exists
	var existingCategory entity.Category
	result := con.First(&existingCategory, "id = $1", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			res.Status = http.StatusNotFound
			res.Message = "Book not found"
			return res, nil
		}
		return res, result.Error
	}

	// Update the Category fields
	existingCategory.Name = payload.Name
	existingCategory.Status = payload.Status
	existingCategory.UpdatedAt = time.Now()

	// Save the updated Category to the database
	if err := con.Save(&existingCategory).Error; err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = existingCategory

	return res, nil
}

func DeleteCategory(id string) (Response, error) {
	var (
		res Response
		con = db.CreateCon()
	)

	// Check if the Category with the given ID exists
	var existingCategory entity.Category
	result := con.First(&existingCategory, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			res.Status = http.StatusNotFound
			res.Message = "Book not found"
			return res, nil
		}
		return res, result.Error
	}

	// Delete the Category from the database
	if err := con.Delete(&existingCategory).Error; err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": 1, // Since you're deleting a single record
	}

	return res, nil
}
