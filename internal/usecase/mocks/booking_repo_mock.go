package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/yousefggg/tour-service/internal/domain"
)

type BookingRepoMock struct {
	CreateBookingFunc   func(ctx context.Context, b *domain.Booking) error
	GetUserBookingsFunc func(ctx context.Context, userID uuid.UUID) ([]domain.Booking, error)
	GetBookingByIDFunc  func(ctx context.Context, id uuid.UUID) (*domain.Booking, error)
}
func (m *BookingRepoMock) CreateBooking(ctx context.Context, b *domain.Booking) error {
	if m.CreateBookingFunc == nil {
		return nil
	}
	return m.CreateBookingFunc(ctx, b)
}
func (m *BookingRepoMock) GetUserBookings(ctx context.Context, userID uuid.UUID) ([]domain.Booking, error) {
	if m.GetUserBookingsFunc == nil {
		return nil, nil
	}
	return m.GetUserBookingsFunc(ctx, userID)
}
func (m *BookingRepoMock) GetBookingByID(ctx context.Context, id uuid.UUID) (*domain.Booking, error) {
	if m.GetBookingByIDFunc == nil {
		return nil, nil
	}
	return m.GetBookingByIDFunc(ctx, id)
}