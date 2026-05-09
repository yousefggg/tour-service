package domain

import (
	"context"

	"github.com/google/uuid"
)
type TourRepository interface {
	Create(ctx context.Context, tour *Tour) error
	GetByID(ctx context.Context, id uuid.UUID) (*Tour, error)
	GetAll(ctx context.Context) ([]*Tour, error)
	Update(ctx context.Context, tour *Tour) error
	Delete(ctx context.Context, id uuid.UUID) error
}
type BookingRepository interface {
	Create(ctx context.Context, booking *Booking) error
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*Booking, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Booking, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status BookingStatus) error
}