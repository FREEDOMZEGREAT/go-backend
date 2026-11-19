package menu

import "sync"

// MemoryStore is a thread-safe in-memory implementation of a dish store.
// It satisfies the DishReader interface declared in the api package.
type MemoryStore struct {
	mu     sync.RWMutex
	dishes map[string]Dish
}

// NewMemoryStore initializes the store with the seed dishes.
func NewMemoryStore() *MemoryStore {
	seeds := seedDishes()
	store := &MemoryStore{
		dishes: make(map[string]Dish, len(seeds)),
	}
	for _, d := range seeds {
		store.dishes[d.ID] = d
	}
	return store
}

// All returns every dish in the store.
func (s *MemoryStore) All() ([]Dish, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]Dish, 0, len(s.dishes))
	for _, d := range s.dishes {
		result = append(result, d)
	}
	return result, nil
}

// ByID returns the dish with the given ID, or ErrNotFound.
func (s *MemoryStore) ByID(id string) (Dish, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	d, ok := s.dishes[id]
	if !ok {
		return Dish{}, ErrNotFound
	}
	return d, nil
}