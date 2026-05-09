package domain

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	TourID    uuid.UUID

	PhoneNumber string

	PeopleCount int

	Notes string

	MedicalInfo string 

	PaymentMethod PaymentMethod

	Status    BookingStatus
	CreatedAt time.Time
}
type PaymentMethod string

const (
	PaymentCard PaymentMethod = "card"
	PaymentCash PaymentMethod = "cash"
	PaymentOnline PaymentMethod = "online"
)
type BookingStatus string

const (
	BookingPending  BookingStatus = "pending"
	BookingApproved BookingStatus = "approved"
	BookingRejected BookingStatus = "rejected"
)