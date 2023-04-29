package invoices

import (
	"database/sql"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

type Repository interface {
	Create(invoices *domain.Invoices) (int64, error)
	ReadAll() ([]*domain.Invoices, error)
	UpdateAll() error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(invoices *domain.Invoices) (int64, error) {
	query := `INSERT INTO invoices (customer_id, datetime, total) VALUES (?, ?, ?)`
	row, err := r.db.Exec(query, &invoices.CustomerId, &invoices.Datetime, &invoices.Total)
	if err != nil {
		return 0, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) ReadAll() ([]*domain.Invoices, error) {
	query := `SELECT id, customer_id, datetime, total FROM invoices`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	invoices := make([]*domain.Invoices, 0)
	for rows.Next() {
		invoice := domain.Invoices{}
		err := rows.Scan(&invoice.Id, &invoice.CustomerId, &invoice.Datetime, &invoice.Total)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, &invoice)
	}
	return invoices, nil
}

func (r *repository) UpdateAll() error {
	query := `UPDATE invoices 
	INNER JOIN (
		SELECT s.invoice_id, SUM(s.quantity * p.price) AS total 
		FROM sales AS s 
		INNER JOIN products AS p ON p.id = s.product_id 
		GROUP BY s.invoice_id
	) AS ventasXfact ON invoices.id = ventasXfact.invoice_id
	SET invoices.total = ventasXfact.total`
	rows, err := r.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
