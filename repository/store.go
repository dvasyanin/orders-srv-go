package repository

import "github.com/go-pg/pg/v10"

type OrderStore interface {
	Offline() OfflineOrdersRepository
	Online() OnlineOrdersRepository
}

// store...
type store struct {
	db                *pg.DB
	onlineRepository  *onlineRepository
	offlineRepository *offlineRepository
}

// NewStore...
func NewStore(db *pg.DB) *store {
	return &store{
		db: db,
	}
}

// Offline...
func (s *store) Offline() OfflineOrdersRepository {
	if s.offlineRepository != nil {
		return s.offlineRepository
	}

	s.offlineRepository = &offlineRepository{
		store: s,
	}

	return s.offlineRepository
}

// Online...
func (s *store) Online() OnlineOrdersRepository {
	if s.onlineRepository != nil {
		return s.onlineRepository
	}

	s.onlineRepository = &onlineRepository{
		store: s,
	}

	return s.onlineRepository
}
