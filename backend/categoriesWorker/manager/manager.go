package manager

import (
	"categoriesWorker/db"
	"protos/parser"
)

type Manager struct {
	parsClient   parser.CategoryParserClient
	db           *db.Database
	workpoolSize int
}

func NewManager(parsClient parser.CategoryParserClient, db *db.Database, workpoolSize int) *Manager {
	return &Manager{parsClient: parsClient, db: db, workpoolSize: workpoolSize}
}
