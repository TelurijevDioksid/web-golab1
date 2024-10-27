package database

import (
	"qrgo/platform/models"

	_ "github.com/lib/pq"
)

func (s *PostgresStorage) GetTotalTickets() (int, error) {
	query := "SELECT COUNT(*) FROM tickets"
	count := 0
	s.db.QueryRow(query).Scan(&count)
	return count, nil
}

func (s *PostgresStorage) GetTicket(id string) (*models.TicketDto, error) {
	query := "SELECT oib, firstname, lastname, createdat FROM tickets WHERE id = $1"
	ticket := new(models.TicketDto)
	s.db.QueryRow(query, id).Scan(
		&ticket.Oib,
		&ticket.FirstName,
		&ticket.LastName,
		&ticket.CreatedAt,
	)
	return ticket, nil
}

func (s *PostgresStorage) GetTotalTicketsByOib(oib string) (int, error) {
	query := "SELECT COUNT(*) FROM tickets WHERE oib = $1"
	count := 0
	s.db.QueryRow(query, oib).Scan(&count)
	return count, nil
}

func (s *PostgresStorage) CreateTicket(dto *models.CreateTicketDto) (string, error) {
	query := "INSERT INTO tickets (oib, firstName, lastName) VALUES ($1, $2, $3) RETURNING id"
	id := ""
	s.db.QueryRow(query, dto.Vatin, dto.FirstName, dto.LastName).Scan(&id)
	return id, nil
}
