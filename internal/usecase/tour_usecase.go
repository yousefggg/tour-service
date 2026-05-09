package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/yousefggg/tour-service/internal/domain"
)
type TourRepository interface {
	GetAllTours(ctx context.Context) ([]domain.Tour, error)
	GetTourByID(ctx context.Context, id uuid.UUID) (*domain.Tour, error)

	CreateTour(ctx context.Context, t *domain.Tour) error
	UpdateTour(ctx context.Context, t *domain.Tour) error
	DeleteTour(ctx context.Context, id uuid.UUID) error
}
type TourUsecase struct {
	repo TourRepository
}

func NewTourUsecase(repo TourRepository) *TourUsecase {
	return &TourUsecase{repo: repo}
}
func (u *TourUsecase) GetAllTours(ctx context.Context) ([]domain.Tour, error) {
	return u.repo.GetAllTours(ctx)
}
func (u *TourUsecase) GetTourByID(ctx context.Context, id uuid.UUID) (*domain.Tour, error) {
	return u.repo.GetTourByID(ctx, id)
}
func (u *TourUsecase) CreateTour(ctx context.Context, t *domain.Tour, role string) error {

	if role != "admin" {
		return domain.ErrForbidden
	}

	if t.Title == "" || t.Price <= 0 {
		return domain.ErrInvalidInput
	}

	t.ID = uuid.New()

	return u.repo.CreateTour(ctx, t)
}
func (u *TourUsecase) UpdateTour(ctx context.Context, t *domain.Tour, role string) error {

	if role != "admin" {
		return domain.ErrForbidden
	}

	if t.ID == uuid.Nil {
		return domain.ErrInvalidInput
	}

	return u.repo.UpdateTour(ctx, t)
}
func (u *TourUsecase) DeleteTour(ctx context.Context, id uuid.UUID, role string) error {

	if role != "admin" {
		return domain.ErrForbidden
	}

	if id == uuid.Nil {
		return domain.ErrInvalidInput
	}

	return u.repo.DeleteTour(ctx, id)
}