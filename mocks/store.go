package mocks

import (
	"orders-srv-go/repository"
	"testing"
)

// store...
type store struct {
	onlineRepository  *onlineRepository
	offlineRepository *offlineRepository
	t                 *testing.T
}

// NewStore...
func NewStore(t *testing.T) *store {
	t.Helper()
	return &store{
		t: t,
	}
}

// Offline...
func (s *store) Offline() repository.OfflineOrdersRepository {
	s.t.Helper()
	if s.offlineRepository != nil {
		return s.offlineRepository
	}

	s.offlineRepository = &offlineRepository{
		store: s,
	}

	return s.offlineRepository
}

// Online...
func (s *store) Online() repository.OnlineOrdersRepository {
	s.t.Helper()
	if s.onlineRepository != nil {
		return s.onlineRepository
	}

	s.onlineRepository = &onlineRepository{
		store: s,
	}

	return s.onlineRepository
}
