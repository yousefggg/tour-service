package usecase

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/yousefggg/tour-service/internal/domain"
	"github.com/yousefggg/tour-service/internal/usecase/mocks"
)
func TestBookingUsecase_CreateBooking_Success(t *testing.T) {

	tourID := uuid.New()
	userID := uuid.New()

	tourRepo := &mocks.TourRepoMock{
		GetTourByIDFunc: func(ctx context.Context, id uuid.UUID) (*domain.Tour, error) {
			return &domain.Tour{ID: tourID}, nil
		},
	}

	bookingRepo := &mocks.BookingRepoMock{
		CreateBookingFunc: func(ctx context.Context, b *domain.Booking) error {
			// проверяем что данные доходят корректно
			if b.UserID != userID {
				t.Fatal("wrong user id")
			}
			if b.TourID != tourID {
				t.Fatal("wrong tour id")
			}
			if b.PeopleCount <= 0 {
				t.Fatal("invalid people count")
			}
			return nil
		},
	}

	uc := NewBookingUsecase(bookingRepo, tourRepo)

	input := CreateBookingInput{
		TourID:        tourID,
		PhoneNumber:   "123456",
		PeopleCount:   2,
		Notes:         "test note",
		MedicalInfo:   "none",
		PaymentMethod: "card",
	}

	err := uc.CreateBooking(context.Background(), userID, input)

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}
func TestBookingUsecase_CreateBooking_TourNotFound(t *testing.T) {

	tourRepo := &mocks.TourRepoMock{
		GetTourByIDFunc: func(ctx context.Context, id uuid.UUID) (*domain.Tour, error) {
			return nil, nil
		},
	}

	bookingRepo := &mocks.BookingRepoMock{}

	uc := NewBookingUsecase(bookingRepo, tourRepo)

	input := CreateBookingInput{
		TourID:      uuid.New(),
		PeopleCount: 1,
	}

	err := uc.CreateBooking(context.Background(), uuid.New(), input)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
func TestBookingUsecase_CreateBooking_InvalidPeopleCount(t *testing.T) {

	tourRepo := &mocks.TourRepoMock{
		GetTourByIDFunc: func(ctx context.Context, id uuid.UUID) (*domain.Tour, error) {
			return &domain.Tour{ID: id}, nil
		},
	}

	bookingRepo := &mocks.BookingRepoMock{}

	uc := NewBookingUsecase(bookingRepo, tourRepo)

	input := CreateBookingInput{
		TourID:      uuid.New(),
		PeopleCount: 0,
	}

	err := uc.CreateBooking(context.Background(), uuid.New(), input)

	if err == nil {
		t.Fatal("expected validation error")
	}
}
func TestBookingUsecase_GetUserBookings(t *testing.T) {

	userID := uuid.New()

	bookingRepo := &mocks.BookingRepoMock{
		GetUserBookingsFunc: func(ctx context.Context, id uuid.UUID) ([]domain.Booking, error) {
			return []domain.Booking{
				{ID: uuid.New(), UserID: id},
				{ID: uuid.New(), UserID: id},
			}, nil
		},
	}

	uc := NewBookingUsecase(bookingRepo, nil)

	result, err := uc.GetUserBookings(context.Background(), userID)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result) != 2 {
		t.Fatalf("expected 2 bookings, got %d", len(result))
	}
}
func TestBookingUsecase_GetBookingByID(t *testing.T) {

	id := uuid.New()

	bookingRepo := &mocks.BookingRepoMock{
		GetBookingByIDFunc: func(ctx context.Context, bid uuid.UUID) (*domain.Booking, error) {
			return &domain.Booking{ID: id}, nil
		},
	}

	uc := NewBookingUsecase(bookingRepo, nil)

	result, err := uc.GetBookingByID(context.Background(), id)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.ID != id {
		t.Fatal("wrong booking returned")
	}
}