package inmemory

import (
	"catcher/internal/models"
	"context"
	"fmt"
	"sync"
)

type PaymentsRepo struct {
	data map[int]*models.Payment

	mu *sync.RWMutex
}

func (db *PaymentsRepo) Create(ctx context.Context, payment *models.Payment) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[payment.ID] = payment
	return nil
}

func (db *PaymentsRepo) All(ctx context.Context) ([]*models.Payment, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	payments := make([]*models.Payment, 0, len(db.data))
	for _, payment := range db.data {
		payments = append(payments, payment)
	}

	return payments, nil
}

func (db *PaymentsRepo) ByID(ctx context.Context, id int) (*models.Payment, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	payment, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("No payment with id %d", id)
	}

	return payment, nil
}

func (db *PaymentsRepo) Update(ctx context.Context, payment *models.Payment) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[payment.ID] = payment
	return nil
}

func (db *PaymentsRepo) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
