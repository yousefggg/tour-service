package domain

import (
	"time"

	"github.com/google/uuid"
)

type Tour struct {
	ID          uuid.UUID
	Title       string
	Location    string
	Country     string
	Season      string
	Price       int64

	ImageURL    string

	Description string
	Includes    string

	CreatedAt   time.Time
	UpdatedAt   time.Time
}