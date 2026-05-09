package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/go-chi/chi/v5"

	"github.com/yousefggg/tour-service/internal/delivery/dto"
	"github.com/yousefggg/tour-service/internal/usecase"
	"github.com/yousefggg/tour-service/internal/domain"
)
type TourHandler struct {
	uc *usecase.TourUsecase
}

func NewTourHandler(uc *usecase.TourUsecase) *TourHandler {
	return &TourHandler{uc: uc}
}
 // GetAllTours godoc
 // @Summary Get all tours
 // @Tags tours
 // @Produce json
 // @Success 200 {array} dto.TourCardResponse
 // @Router /tours [get]
func (h *TourHandler) GetAllTours(w http.ResponseWriter, r *http.Request) {
	tours, err := h.uc.GetAllTours(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := make([]dto.TourCardResponse, 0, len(tours))
	for _, t := range tours {
		resp = append(resp, dto.TourCardResponse{
			ID:       t.ID,
			Title:    t.Title,
			Location: t.Location,
			Country:  t.Country,
			Season:   t.Season,
			Price:    t.Price,
			ImageURL: t.ImageURL,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
 // GetTourByID godoc
 // @Summary Get tour by ID
 // @Tags tours
 // @Produce json
 // @Param id path string true "Tour ID"
 // @Success 200 {object} dto.TourDetailResponse
 // @Failure 404
 // @Router /tours/{id} [get]
func (h *TourHandler) GetTourByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid uuid", http.StatusBadRequest)
		return
	}

	tour, err := h.uc.GetTourByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := dto.TourDetailResponse{
		ID:          tour.ID,
		Title:       tour.Title,
		Location:    tour.Location,
		Country:     tour.Country,
		Season:      tour.Season,
		Price:       tour.Price,
		ImageURL:    tour.ImageURL,
		Description: tour.Description,
		Includes:    tour.Includes,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
 // CreateTour godoc
 // @Summary Create tour (admin only)
 // @Tags admin-tours
 // @Accept json
 // @Produce json
 // @Param input body dto.CreateTourRequest true "Tour data"
 // @Success 200
 // @Failure 403
 // @Router /admin/tours [post]
func (h *TourHandler) CreateTour(w http.ResponseWriter, r *http.Request) {
	role := r.Header.Get("role")

	var req dto.CreateTourRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	t := &domain.Tour{
		Title:       req.Title,
		Location:    req.Location,
		Country:     req.Country,
		Season:      req.Season,
		Price:       req.Price,
		ImageURL:    req.ImageURL,
		Description: req.Description,
		Includes:    req.Includes,
	}

	err := h.uc.CreateTour(r.Context(), t, role)
	if err != nil {
		if err == domain.ErrForbidden {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"id": t.ID.String(),
	})
}
 // UpdateTour godoc
 // @Summary Update tour (admin only)
 // @Tags admin-tours
 // @Accept json
 // @Produce json
 // @Param id path string true "Tour ID"
 // @Param input body dto.UpdateTourRequest true "Update data"
 // @Success 200
 // @Failure 404
 // @Router /admin/tours/{id} [put]
func (h *TourHandler) UpdateTour(w http.ResponseWriter, r *http.Request) {
	role := r.Header.Get("role")

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid uuid", http.StatusBadRequest)
		return
	}

	var req dto.UpdateTourRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	t := &domain.Tour{
		ID:          id,
		Title:       req.Title,
		Location:    req.Location,
		Country:     req.Country,
		Season:      req.Season,
		Price:       req.Price,
		ImageURL:    req.ImageURL,
		Description: req.Description,
		Includes:    req.Includes,
	}

	err = h.uc.UpdateTour(r.Context(), t, role)
	if err != nil {
		if err == domain.ErrForbidden {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
 // DeleteTour godoc
 // @Summary Delete tour (admin only)
 // @Tags admin-tours
 // @Param id path string true "Tour ID"
 // @Success 200
 // @Failure 404
 // @Router /admin/tours/{id} [delete]
func (h *TourHandler) DeleteTour(w http.ResponseWriter, r *http.Request) {
	role := r.Header.Get("role")

	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid uuid", http.StatusBadRequest)
		return
	}

	err = h.uc.DeleteTour(r.Context(), id, role)
	if err != nil {
		if err == domain.ErrForbidden {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}