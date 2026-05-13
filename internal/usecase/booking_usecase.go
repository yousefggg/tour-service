package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/yousefggg/tour-service/internal/domain"
)

type BookingRepository interface {
	CreateBooking(ctx context.Context, b *domain.Booking) error
	GetUserBookings(ctx context.Context, userID uuid.UUID) ([]domain.Booking, error)
	GetBookingByID(ctx context.Context, id uuid.UUID) (*domain.Booking, error)
}


type BookingUsecase struct {
	bookingRepo BookingRepository
	tourRepo    TourRepository
}



func NewBookingUsecase(
	br BookingRepository,
	tr TourRepository,
) *BookingUsecase {
	return &BookingUsecase{
		bookingRepo: br,
		tourRepo:    tr,
	}
}

type CreateBookingInput struct {
	TourID        uuid.UUID
	PhoneNumber   string
	PeopleCount   int
	Notes         string
	MedicalInfo   string
	PaymentMethod string
}


func (u *BookingUsecase) CreateBooking(
	ctx context.Context,
	userID uuid.UUID,
	input CreateBookingInput,
) error {

	tour, err := u.tourRepo.GetTourByID(ctx, input.TourID)
	if err != nil {
		return err
	}
	if tour == nil {
		return domain.ErrNotFound
	}

	if input.PeopleCount <= 0 {
		return domain.ErrInvalidInput
	}

	booking := &domain.Booking{
		ID:     uuid.New(),
		UserID: userID,
		TourID: input.TourID,

		PhoneNumber:   input.PhoneNumber,
		PeopleCount:   input.PeopleCount,
		Notes:         input.Notes,
		MedicalInfo:   input.MedicalInfo,
		PaymentMethod: domain.PaymentMethod(input.PaymentMethod),

		Status: domain.BookingPending,
	}

	return u.bookingRepo.CreateBooking(ctx, booking)
}



func (u *BookingUsecase) GetUserBookings(
	ctx context.Context,
	userID uuid.UUID,
) ([]domain.Booking, error) {

	return u.bookingRepo.GetUserBookings(ctx, userID)
}


func (u *BookingUsecase) GetBookingByID(
	ctx context.Context,
	id uuid.UUID,
) (*domain.Booking, error) {

	return u.bookingRepo.GetBookingByID(ctx, id)
}