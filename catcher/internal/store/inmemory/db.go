package inmemory

import (
	"catcher/internal/models"
	"catcher/internal/store"
	"sync"
)

type DB struct {
	paymentsRepo store.PaymentsRepository

	mu *sync.RWMutex
}

func NewDB() store.Store {
	return &DB{
		mu: new(sync.RWMutex),
	}
}

func (db *DB) Payments() store.PaymentsRepository {
	if db.paymentsRepo == nil {
		db.paymentsRepo = &PaymentsRepo{
			data: make(map[int]*models.Payment),
			mu:   new(sync.RWMutex),
		}
	}

	return db.paymentsRepo
}
