// handlers/opera_handler.go

package handlers

import (
	"context"
	"libretto/models"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
)

var db *pgxpool.Pool

func SetupRoutes(e *echo.Echo, pool *pgxpool.Pool) {
	db = pool
	e.POST("/operas", createOpera)
	// e.DELETE("/operas/:id", deleteOpera)
	// e.GET("/operas", findOpera)
	// e.PUT("/operas/:id", editOpera)
}

func createOpera(c echo.Context) error {
	opera := new(models.Opera)
	if err := c.Bind(opera); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload", "details": err.Error()})
	}

	query := "INSERT INTO operas (title, composer, premiere, description, language, genre, librettist) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	var operaID int
	err := db.QueryRow(context.Background(), query, opera.Title, opera.Composer, opera.Premiere, opera.Description, opera.Language, opera.Genre, opera.Librettist).Scan(&operaID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	opera.ID = operaID
	return c.JSON(http.StatusCreated, opera)
}
