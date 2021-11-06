package store

import (
	"catcher/internal/models"
	"context"
)

type Store interface {
	Payments() PaymentsRepository
}

type PaymentsRepository interface {
	Create(ctx context.Context, payment *models.Payment) error
	All(ctx context.Context) ([]*models.Payment, error)
	ByID(ctx context.Context, id int) (*models.Payment, error)
	Update(ctx context.Context, payment *models.Payment) error
	Delete(ctx context.Context, id int) error
}
