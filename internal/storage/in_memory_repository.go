package storage

import (
	"math"
	"math/rand"
	"sync"
)

type InMemoryFactsRepository struct {
	store      []string
	storeMutex sync.RWMutex
}

func InitInMemoryFactsRepository() *InMemoryFactsRepository {
	return &InMemoryFactsRepository{
		store:      make([]string, 0, math.MaxUint16),
		storeMutex: sync.RWMutex{},
	}
}

func (r *InMemoryFactsRepository) SaveFact(fact string) error {
	r.storeMutex.Lock()
	r.store = append(r.store, fact)
	r.storeMutex.Unlock()
	return nil
}

func (r *InMemoryFactsRepository) GetFact() (string, error) {
	return r.store[rand.Intn(len(r.store))], nil
}
