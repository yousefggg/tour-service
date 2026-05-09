package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/yousefggg/tour-service/internal/delivery/dto"
	"github.com/yousefggg/tour-service/internal/domain"
	"github.com/yousefggg/tour-service/internal/usecase"
)

type BookingHandler struct {
	uc *usecase.BookingUsecase
}

func NewBookingHandler(uc *usecase.BookingUsecase) *BookingHandler {
	return &BookingHandler{uc: uc}
}

func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	userID, err := h.getUserID(r)
	if err != nil {
		h.respondWithError(w, http.StatusUnauthorized, "invalid user id")
		return
	}

	var req dto.CreateBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	tourID, err := uuid.Parse(req.TourID)
	if err != nil {
		h.respondWithError(w, http.StatusBadRequest, "invalid tour id format")
		return
	}

	err = h.uc.CreateBooking(
		r.Context(),
		userID,
		usecase.CreateBookingInput{
			TourID:        tourID,
			PhoneNumber:   req.PhoneNumber,
			PeopleCount:   req.PeopleCount,
			Notes:         req.Notes,
			MedicalInfo:   req.MedicalInfo,
			PaymentMethod: req.PaymentMethod,
		},
	)

	if err != nil {
		if err == domain.ErrNotFound {
			h.respondWithError(w, http.StatusNotFound, "tour not found")
			return
		}
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondWithJSON(w, http.StatusCreated, map[string]string{
		"status": "booking created",
	})
}

// --- Get User Bookings ---
func (h *BookingHandler) GetUserBookings(w http.ResponseWriter, r *http.Request) {
	userID, err := h.getUserID(r)
	if err != nil {
		h.respondWithError(w, http.StatusUnauthorized, "invalid user id")
		return
	}

	bookings, err := h.uc.GetUserBookings(r.Context(), userID)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := make([]dto.BookingResponse, 0, len(bookings))
	for _, b := range bookings {
		resp = append(resp, mapToResponse(b))
	}

	h.respondWithJSON(w, http.StatusOK, resp)
}

func (h *BookingHandler) GetBookingByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		h.respondWithError(w, http.StatusBadRequest, "invalid booking uuid")
		return
	}

	booking, err := h.uc.GetBookingByID(r.Context(), id)
	if err != nil {
		h.respondWithError(w, http.StatusNotFound, "booking not found")
		return
	}

	h.respondWithJSON(w, http.StatusOK, mapToResponse(*booking))
}

func (h *BookingHandler) getUserID(r *http.Request) (uuid.UUID, error) {
	return uuid.Parse(r.Header.Get("user_id"))
}

func (h *BookingHandler) respondWithError(w http.ResponseWriter, code int, message string) {
	h.respondWithJSON(w, code, map[string]string{"error": message})
}

func (h *BookingHandler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func mapToResponse(b domain.Booking) dto.BookingResponse {
	return dto.BookingResponse{
		ID:            b.ID,
		TourID:        b.TourID,
		TourTitle:     "",
		Price:         0,
		Status:        string(b.Status),
		PhoneNumber:   b.PhoneNumber,
		PeopleCount:   b.PeopleCount,
		Notes:         b.Notes,
		MedicalInfo:   b.MedicalInfo,
		PaymentMethod: string(b.PaymentMethod),
		CreatedAt:     b.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}