package models

import (
	"time"
)

type Opera struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Composer    string    `json:"composer"`
	Premiere    time.Time `json:"premiere"`
	Description string    `json:"description"`
	Language    string    `json:"language"`
	Genre       string    `json:"genre"`
	Librettist  string    `json:"librettist"`
}
