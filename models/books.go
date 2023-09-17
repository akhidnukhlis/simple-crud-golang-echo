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

func FetchAllBook() (Response, error) {
	var (
		books entity.Books
		res   Response

		con = db.CreateCon()
	)

	sqlStatement := "SELECT * FROM books"

	con.Raw(sqlStatement).Scan(&books)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = books

	return res, nil
}

func GetByIdBook(id string) (*Response, error) {
	var (
		book entity.Book
		res  Response

		con = db.CreateCon()
	)

	// Assuming "con" is your GORM database connection, and "id" is the ID you want to query.
	err := con.Where("id = $1", id).First(&book).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Handle the case when the book is not found
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
	res.Data = book

	return &res, nil
}

func StoreBook(payload *entity.BookRequest) (*Response, error) {
	var (
		res Response
		con = db.CreateCon()
	)

	// Create a new Book instance with the payload data
	book := &entity.Book{
		ID:          uuid.New().String(),
		Tittle:      payload.Tittle,
		Description: payload.Description,
		Price:       payload.Price,
		Image:       payload.Image,
		Categories:  payload.Categories,
		Keywords:    payload.Keywords,
		Stock:       payload.Stock,
		Publisher:   payload.Publisher,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Insert the new Book record into the database
	if err := con.Create(book).Error; err != nil {
		err = fmt.Errorf("creating book got: %w", err)

		return nil, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = book

	return &res, nil
}

func UpdateBook(id string, payload *entity.BookRequest) (Response, error) {
	var (
		res Response
		con = db.CreateCon()
	)

	// Check if the book with the given ID exists
	var existingBook entity.Book
	result := con.First(&existingBook, "id = $1", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			res.Status = http.StatusNotFound
			res.Message = "Book not found"
			return res, nil
		}
		return res, result.Error
	}

	// Update the book fields
	existingBook.Tittle = payload.Tittle
	existingBook.Description = payload.Description
	existingBook.Price = payload.Price
	existingBook.Image = payload.Image
	existingBook.Categories = payload.Categories
	existingBook.Keywords = payload.Keywords
	existingBook.Stock = payload.Stock
	existingBook.Publisher = payload.Publisher
	existingBook.UpdatedAt = time.Now()

	// Save the updated book to the database
	if err := con.Save(&existingBook).Error; err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = existingBook

	return res, nil
}

func DeleteBook(id string) (Response, error) {
	var (
		res Response
		con = db.CreateCon()
	)

	// Check if the book with the given ID exists
	var existingBook entity.Book
	result := con.First(&existingBook, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			res.Status = http.StatusNotFound
			res.Message = "Book not found"
			return res, nil
		}
		return res, result.Error
	}

	// Delete the book from the database
	if err := con.Delete(&existingBook).Error; err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": 1, // Since you're deleting a single record
	}

	return res, nil
}
