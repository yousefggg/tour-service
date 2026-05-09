package dto

import "github.com/google/uuid"

type CreateBookingRequest struct {
	TourID        string `json:"tour_id"`
	PhoneNumber   string `json:"phone_number"`
	PeopleCount   int    `json:"people_count"`
	Notes         string `json:"notes"`
	MedicalInfo   string `json:"medical_info"`
	PaymentMethod string `json:"payment_method"`
}

type BookingResponse struct {
	ID            uuid.UUID `json:"id"`
	TourID        uuid.UUID `json:"tour_id"`
	TourTitle     string    `json:"tour_title"`
	Price         int64     `json:"price"`
	Status        string    `json:"status"`
	PhoneNumber   string    `json:"phone_number"`
	PeopleCount   int       `json:"people_count"`
	Notes         string    `json:"notes"`
	MedicalInfo   string    `json:"medical_info"`
	PaymentMethod string    `json:"payment_method"`
	CreatedAt     string    `json:"created_at"`
}