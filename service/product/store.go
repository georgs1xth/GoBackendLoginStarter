package product

import (
	"database/sql"
	"fmt"
	"log/slog"
	"strings"

	"github.com/georgs1xth/APIBACKEND/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}
	return products, nil
}

func (s *Store) GetProductsByIDs(productIDs []int) ([]types.Product, error) {
	placeholders := make([]string, len(productIDs))
	for i := range productIDs {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (%s)", strings.Join(placeholders, ", "))

	// Convert productIDs to []interface{}
	args := make([]interface{}, len(productIDs))
	for i, v := range productIDs {
		args[i] = v
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	products := []types.Product{}
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) UpdateProduct(product types.Product) error {
	_, err := s.db.Exec("UPDATE products SET name = $1, price = $2, iamge = $3, description = $4, quantity = $5 WHERE id = $6", product.Name, product.Price, product.Image, product.Description, product.Quantity, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)
	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Store) CreateProduct(product types.Product) error {
	_, err := s.db.Exec("INSERT INTO products (name, description, image, price, quantity) VALUES ($1, $2, $3, $4, $5)", product.Name, product.Description, product.Image, product.Price, product.Quantity)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	slog.Info("product creation", "product", product.Name)
	return nil
}
