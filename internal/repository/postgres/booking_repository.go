package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/yousefggg/tour-service/internal/domain"
	"github.com/yousefggg/common-lib/pkg/logger"
	"github.com/yousefggg/common-lib/pkg/errors"
)

type BookingRepository struct {
	db *pgxpool.Pool
}

func NewBookingRepository(db *pgxpool.Pool) *BookingRepository {
	return &BookingRepository{db: db}
}
 // CreateBooking godoc
 // @Summary Create booking
 // @Description Create new tour booking with user details
 // @Tags bookings
 // @Security BearerAuth
 // @Accept json
 // @Produce json
 // @Param input body dto.CreateBookingRequest true "Booking request"
 // @Success 201 {object} map[string]string
 // @Failure 400
 // @Failure 401
 // @Failure 500
 // @Router /bookings [post]
func (r *BookingRepository) CreateBooking(ctx context.Context, b *domain.Booking) error {

	query := `
		INSERT INTO bookings (
			user_id,
			tour_id,
			phone_number,
			people_count,
			notes,
			medical_info,
			payment_method,
			status,
			created_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,NOW())
		RETURNING id, created_at
	`

	err := r.db.QueryRow(
		ctx,
		query,
		b.UserID,
		b.TourID,
		b.PhoneNumber,
		b.PeopleCount,
		b.Notes,
		b.MedicalInfo,
		b.PaymentMethod,
		b.Status,
	).Scan(&b.ID, &b.CreatedAt)

	if err != nil {
		logger.Error("failed to create booking", "error", err)
		return errors.NewErr("DB_CREATE_BOOKING_FAILED", "failed to create booking", err)
	}

	return nil
}
 // GetUserBookings godoc
 // @Summary Get user bookings
 // @Description Get all bookings for current user
 // @Tags bookings
 // @Security BearerAuth
 // @Produce json
 // @Success 200 {array} dto.BookingResponse
 // @Failure 401
 // @Failure 500
 // @Router /bookings [get]
func (r *BookingRepository) GetUserBookings(ctx context.Context, userID uuid.UUID) ([]domain.Booking, error) {

	query := `
		SELECT 
			id,
			user_id,
			tour_id,
			phone_number,
			people_count,
			notes,
			medical_info,
			payment_method,
			status,
			created_at
		FROM bookings
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		logger.Error("failed to get user bookings", "error", err, "user_id", userID)
		return nil, errors.NewErr("DB_GET_BOOKINGS_FAILED", "failed to get bookings", err)
	}
	defer rows.Close()

	var bookings []domain.Booking

	for rows.Next() {
		var b domain.Booking

		err := rows.Scan(
			&b.ID,
			&b.UserID,
			&b.TourID,
			&b.PhoneNumber,
			&b.PeopleCount,
			&b.Notes,
			&b.MedicalInfo,
			&b.PaymentMethod,
			&b.Status,
			&b.CreatedAt,
		)

		if err != nil {
			logger.Error("failed to scan booking", "error", err)
			return nil, err
		}

		bookings = append(bookings, b)
	}

	return bookings, nil
}
 // GetBookingByID godoc
 // @Summary Get booking by ID
 // @Description Get single booking by UUID
 // @Tags bookings
 // @Security BearerAuth
 // @Produce json
 // @Param id path string true "Booking ID"
 // @Success 200 {object} dto.BookingResponse
 // @Failure 400
 // @Failure 404
 // @Router /bookings/{id} [get]
func (r *BookingRepository) GetBookingByID(ctx context.Context, id uuid.UUID) (*domain.Booking, error) {

	query := `
		SELECT 
			id,
			user_id,
			tour_id,
			phone_number,
			people_count,
			notes,
			medical_info,
			payment_method,
			status,
			created_at
		FROM bookings
		WHERE id = $1
	`

	var b domain.Booking

	err := r.db.QueryRow(ctx, query, id).Scan(
		&b.ID,
		&b.UserID,
		&b.TourID,
		&b.PhoneNumber,
		&b.PeopleCount,
		&b.Notes,
		&b.MedicalInfo,
		&b.PaymentMethod,
		&b.Status,
		&b.CreatedAt,
	)

	if err != nil {
		logger.Error("failed to get booking", "error", err, "booking_id", id)
		return nil, errors.NewErr("BOOKING_NOT_FOUND", "booking not found", err)
	}

	return &b, nil
}