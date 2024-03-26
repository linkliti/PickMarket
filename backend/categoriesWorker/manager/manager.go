package manager

import (
	"categoriesWorker/db"
	"protos/parser"
)

type Manager struct {
	parsClient parser.CategoryParserClient
	db         *db.Database
}

func NewManager(parsClient parser.CategoryParserClient, db *db.Database) *Manager {
	return &Manager{parsClient: parsClient, db: db}
}
