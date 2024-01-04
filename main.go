package main

import (
	"context"

	"libretto/handlers"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	pool, err := pgxpool.Connect(context.Background(), "postgres://pedro:limonada@localhost/libretto_press?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	e := echo.New()
	handlers.SetupRoutes(e, pool)

	e.Logger.Fatal(e.Start(":8080"))
}
