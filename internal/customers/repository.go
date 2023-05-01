package customers

import (
	"database/sql"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

type Top5products struct {
	Description string `json:"desc"`
	Count       int    `json:"count"`
}

type Repository interface {
	Create(customers *domain.Customers) (int64, error)
	ReadAll() ([]*domain.Customers, error)
	GetTop5SoldPrd() ([]*Top5products, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(customers *domain.Customers) (int64, error) {
	query := `INSERT INTO customers (customers.first_name, customers.last_name, customers.condition) VALUES ( ?, ?, ?)`
	row, err := r.db.Exec(query, &customers.FirstName, &customers.LastName, &customers.Condition)
	if err != nil {
		return 0, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) ReadAll() ([]*domain.Customers, error) {
	query := `SELECT id, first_name, last_name, customers.condition FROM customers`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	customers := make([]*domain.Customers, 0)
	for rows.Next() {
		customer := domain.Customers{}
		err := rows.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Condition)
		if err != nil {
			return nil, err
		}
		customers = append(customers, &customer)
	}
	return customers, nil
}

func (r *repository) GetTop5SoldPrd() ([]*Top5products, error) {
	query := `select products.description, sum(sales.quantity) as ct from sales inner join products on sales.product_id = products.id
	group by sales.product_id
	order by ct desc
	LIMIT 0, 3`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]*Top5products, 0)
	for rows.Next() {
		row := Top5products{}
		err := rows.Scan(&row.Description, &row.Count)
		if err != nil {
			return nil, err
		}
		result = append(result, &row)
	}
	return result, nil
}
