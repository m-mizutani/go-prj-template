package infra

import (
	"github.com/m-mizutani/go-prj-template/pkg/infra/db"
)

type Clients struct {
	db *db.Client
}

func New(options ...Option) *Clients {
	clients := &Clients{}

	for _, opt := range options {
		opt(clients)
	}

	return clients
}

func (x *Clients) DB() *db.Client {
	if x.db == nil {
		panic("infra.Clients.DB is not configured, but called")
	}
	return x.db
}

type Option func(clients *Clients)

func WithDB(dbClient *db.Client) Option {
	return func(clients *Clients) {
		clients.db = dbClient
	}
}
