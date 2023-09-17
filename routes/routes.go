package routes

import (
	"github.com/akhidnukhlis/controllers"
	"github.com/akhidnukhlis/middleware"
	"github.com/labstack/echo"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, to simple crud go-echo!!")
	})

	// api data book
	e.GET("/api/v1/books", controllers.FetchAllBook, middleware.IsAuthenticated)
	e.GET("/api/v1/book/:id", controllers.GetByIdBook, middleware.IsAuthenticated)
	e.POST("/api/v1/book", controllers.StoreBook, middleware.IsAuthenticated)
	e.PUT("/api/v1/book/:id", controllers.UpdateBook, middleware.IsAuthenticated)
	e.DELETE("/api/v1/book/:id", controllers.DeleteBook, middleware.IsAuthenticated)

	// api data category
	e.GET("/api/v1/categories", controllers.FetchAllCategory, middleware.IsAuthenticated)
	e.GET("/api/v1/category/:id", controllers.GetByIdCategory, middleware.IsAuthenticated)
	e.POST("/api/v1/category", controllers.StoreCategory, middleware.IsAuthenticated)
	e.PUT("/api/v1/category/:id", controllers.UpdateCategory, middleware.IsAuthenticated)
	e.DELETE("/api/v1/category/:id", controllers.DeleteCategory, middleware.IsAuthenticated)

	// api data user
	e.POST("/api/v1/user", controllers.StoreUser)

	// api hash & login
	e.GET("/api/v1/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/api/v1/login", controllers.CheckLogin)

	// api unit tests product

	return e
}
