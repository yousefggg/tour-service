package postgres
import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yousefggg/tour-service/internal/domain"
	"github.com/yousefggg/common-lib/pkg/logger"
	"github.com/yousefggg/common-lib/pkg/errors"
	"github.com/google/uuid"
)

type TourRepository struct {
	db *pgxpool.Pool
}

func NewTourRepository(db *pgxpool.Pool) *TourRepository {
	return &TourRepository{
		db: db,
	}
}

func (r *TourRepository) GetAllTours(ctx context.Context) ([]domain.Tour, error) {
	query := `
		SELECT id, title, location, price, country, season, image, description, includes
		FROM tours
		ORDER BY id DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tours []domain.Tour

	for rows.Next() {
		var t domain.Tour

		err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Location,
			&t.Price,
			&t.Country,
			&t.Season,
			&t.ImageURL,
			&t.Description,
			&t.Includes,
		)
		if err != nil {
			return nil, err
		}

		tours = append(tours, t)
	}

	return tours, nil
}

func (r *TourRepository) GetTourByID(ctx context.Context, id uuid.UUID) (*domain.Tour, error) {
	query := `
		SELECT id, title, location, price, country, season, image, description, includes
		FROM tours
		WHERE id = $1
	`

	var t domain.Tour

	err := r.db.QueryRow(ctx, query, id).Scan(
		&t.ID,
		&t.Title,
		&t.Location,
		&t.Price,
		&t.Country,
		&t.Season,
		&t.ImageURL,
		&t.Description,
		&t.Includes,
	)

	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (r *TourRepository) CreateTour(ctx context.Context, t *domain.Tour) error {
	query := `
		INSERT INTO tours (title, location, price, country, season, image, description, includes)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		t.Title,
		t.Location,
		t.Price,
		t.Country,
		t.Season,
		t.ImageURL,
		t.Description,
		t.Includes,
	).Scan(&t.ID)

	if err != nil {
		return err
	}

	return nil
}
func (r *TourRepository) UpdateTour(ctx context.Context, t *domain.Tour) error {
	query := `
		UPDATE tours
		SET title = $1,
			location = $2,
			price = $3,
			country = $4,
			season = $5,
			image = $6,
			description = $7,
			includes = $8
		WHERE id = $9
	`

	cmdTag, err := r.db.Exec(ctx, query,
		t.Title,
		t.Location,
		t.Price,
		t.Country,
		t.Season,
		t.ImageURL,
		t.Description,
		t.Includes,
		t.ID,
	)

	if err != nil {
		logger.Error("failed to update tour", "error", err, "tour_id", t.ID)
		return errors.NewErr("DB_UPDATE_FAILED", "failed to update tour", err)
	}

	if cmdTag.RowsAffected() == 0 {
		logger.Warn("tour not found on update", "tour_id", t.ID)
		return errors.NewErr("TOUR_NOT_FOUND", "tour not found", nil)
	}

	return nil
}
func (r *TourRepository) DeleteTour(ctx context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM tours
		WHERE id = $1
	`

	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		logger.Error("failed to delete tour", "error", err, "tour_id", id)
		return errors.NewErr("DB_DELETE_FAILED", "failed to delete tour", err)
	}

	if cmdTag.RowsAffected() == 0 {
		logger.Warn("tour not found on delete", "tour_id", id)
		return errors.NewErr("TOUR_NOT_FOUND", "tour not found", nil)
	}

	return nil
}