package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/yousefggg/tour-service/internal/domain"
)


type TourRepoMock struct {
	GetTourByIDFunc func(ctx context.Context, id uuid.UUID) (*domain.Tour, error)

	CreateTourFunc func(ctx context.Context, t *domain.Tour) error

	GetAllToursFunc func(ctx context.Context) ([]domain.Tour, error)

	UpdateTourFunc func(ctx context.Context, t *domain.Tour) error

	DeleteTourFunc func(ctx context.Context, id uuid.UUID) error
}

func (m *TourRepoMock) GetTourByID(ctx context.Context, id uuid.UUID) (*domain.Tour, error) {
	return m.GetTourByIDFunc(ctx, id)
}

func (m *TourRepoMock) CreateTour(ctx context.Context, t *domain.Tour) error {
	return m.CreateTourFunc(ctx, t)
}

func (m *TourRepoMock) GetAllTours(ctx context.Context) ([]domain.Tour, error) {
	return m.GetAllToursFunc(ctx)
}

func (m *TourRepoMock) UpdateTour(ctx context.Context, t *domain.Tour) error {
	return m.UpdateTourFunc(ctx, t)
}

func (m *TourRepoMock) DeleteTour(ctx context.Context, id uuid.UUID) error {
	return m.DeleteTourFunc(ctx, id)
}