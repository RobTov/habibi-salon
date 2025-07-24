package store

import (
	"context"
	"database/sql"
)

type Stock struct {
	ID        int64 `json:"id"`
	ProductID int64 `json:"product_id"`
	ServiceID int64 `json:"service_id"`
	Quantity  int64 `json:"quantity"`
}

type StockStore struct {
	db *sql.DB
}

func (s *StockStore) GetAll(ctx context.Context) ([]Stock, error) {
	query := `
	SELECT id, product_id, service_id, quantity
	FROM stock;
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	stocks := []Stock{}
	for rows.Next() {
		var s = Stock{}

		err := rows.Scan(
			&s.ID,
			&s.ProductID,
			&s.ServiceID,
			&s.Quantity,
		)

		if err != nil {
			return nil, err
		}

		stocks = append(stocks, s)
	}

	return stocks, nil
}

func (s *StockStore) Create(ctx context.Context, stock *Stock) error {
	query := `
	INSERT INTO stock (product_id, service_id, quantity)
	VALUES ($1, $2, $3);
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, stock.ProductID, stock.ServiceID, stock.Quantity)
	if err != nil {
		return err
	}

	return nil
}
