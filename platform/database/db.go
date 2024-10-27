package database

import (
	"database/sql"
	"qrgo/platform/models"

	_ "github.com/lib/pq"
)

type Storage interface {
	GetTotalTickets() (int, error)
	GetTicket(string) (*models.TicketDto, error)
	GetTotalTicketsByOib(string) (int, error)
	CreateTicket(*models.CreateTicketDto) (string, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func New(conStr string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) Setup() error {
	query := `
        CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

        CREATE TABLE IF NOT EXISTS tickets (
            id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
            oib VARCHAR(100) NOT NULL,
            firstName VARCHAR(100) NOT NULL,
            lastName VARCHAR(100) NOT NULL,
            createdAt TIMESTAMP NOT NULL DEFAULT NOW()
        );
    `
	_, err := s.db.Exec(query)
	return err
}
