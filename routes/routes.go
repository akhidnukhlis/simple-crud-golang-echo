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

	// api data product
	e.GET("/api/product", controllers.FetchAllProduct, middleware.IsAuthenticated)
	e.POST("/api/product", controllers.StoreProduct, middleware.IsAuthenticated)
	e.PUT("/api/product", controllers.UpdateProduct, middleware.IsAuthenticated)
	e.DELETE("/api/product", controllers.DeleteProduct, middleware.IsAuthenticated)

	// api hash & login
	e.GET("/api/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/api/login", controllers.CheckLogin)

	// api unit tests
	e.GET("/api/tests-struct-validation", controllers.TestStructValidation)
	e.GET("/api/tests-variable-validation", controllers.TestVariableValidation)

	// api unit tests product

	return e
}
