package order

import (
	"database/sql"

	"github.com/georgs1xth/APIBACKEND/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateOrder(order types.Order) (int, error) {
	_, err := s.db.Exec("INSERT INTO orders (user_id, total, status, address) VALUES ($1, $2, $3, $4)", order.UserID, order.Total, order.Status, order.Address)
	if err != nil {
		return 0, err
	}
	rows, err := s.db.Query("SELECT * FROM orders WHERE user_id = $1 ORDER BY id DESC LIMIT 1", order.UserID)
	if err != nil {
		return 0, err
	}

	orderSql, err := scanRowsIntoOrder(rows)
	if err != nil {
		return 0, err
	}
	id := orderSql.ID
	return int(id), err
}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	_, err := s.db.Exec("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)", orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	return err
}

func scanRowsIntoOrder(rows *sql.Rows) (*types.Order, error) {
	rows.Next()
	order := new(types.Order)
	err := rows.Scan(
		&order.ID,
		&order.UserID,
		&order.Total,
		&order.Status,
		&order.Address,
		&order.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return order, nil
}
