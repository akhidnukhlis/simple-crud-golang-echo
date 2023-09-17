package main

import (
	"github.com/akhidnukhlis/db"
	"github.com/akhidnukhlis/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8081"))
}
