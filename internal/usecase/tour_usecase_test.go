package usecase

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/yousefggg/tour-service/internal/domain"
	"github.com/yousefggg/tour-service/internal/usecase/mocks"
)
func TestTourUsecase_CreateTour_Success(t *testing.T) {

	tourRepo := &mocks.TourRepoMock{
		CreateTourFunc: func(ctx context.Context, tour *domain.Tour) error {
			if tour.Title == "" {
				t.Fatal("title is empty")
			}
			if tour.Price <= 0 {
				t.Fatal("invalid price")
			}
			return nil
		},
	}

	uc := NewTourUsecase(tourRepo)

	tour := &domain.Tour{
		Title: "Everest Base Camp",
		Price: 1200,
	}

	// 🔥 FIX: добавляем 3-й аргумент (например creator/admin/userId)
	err := uc.CreateTour(context.Background(), tour, "admin")

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}
func TestTourUsecase_GetAllTours(t *testing.T) {

	tourRepo := &mocks.TourRepoMock{
		GetAllToursFunc: func(ctx context.Context) ([]domain.Tour, error) {
			return []domain.Tour{
				{ID: uuid.New(), Title: "Tour 1"},
				{ID: uuid.New(), Title: "Tour 2"},
			}, nil
		},
	}

	uc := NewTourUsecase(tourRepo)

	res, err := uc.GetAllTours(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(res) != 2 {
		t.Fatalf("expected 2 tours, got %d", len(res))
	}
}
func TestTourUsecase_GetTourByID_Success(t *testing.T) {

	id := uuid.New()

	tourRepo := &mocks.TourRepoMock{
		GetTourByIDFunc: func(ctx context.Context, tid uuid.UUID) (*domain.Tour, error) {
			return &domain.Tour{
				ID:    id,
				Title: "Test Tour",
			}, nil
		},
	}

	uc := NewTourUsecase(tourRepo)

	res, err := uc.GetTourByID(context.Background(), id)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.ID != id {
		t.Fatal("wrong tour returned")
	}
}
func TestTourUsecase_DeleteTour_Success(t *testing.T) {

	id := uuid.New()

	tourRepo := &mocks.TourRepoMock{
		DeleteTourFunc: func(ctx context.Context, tid uuid.UUID) error {
			return nil
		},
	}

	uc := NewTourUsecase(tourRepo)

	err := uc.DeleteTour(context.Background(), id, "user")

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}