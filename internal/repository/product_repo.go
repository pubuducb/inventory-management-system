package repository

import (
	"database/sql"
	"ims/internal/model"
)

type ProductRepository interface {
	Create(product *model.Product) error
	GetByID(id int) (*model.Product, error)
	Update(product *model.Product) error
	Archive(id int) error
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (repo *productRepository) Create(product *model.Product) error {
	query := `
		INSERT INTO products (name, price)
		VALUES (?, ?)
		RETURNING id, created_at
	`
	err := repo.db.QueryRow(query, product.Name, product.Price).Scan(&product.ID, &product.CreatedAt)
	return err
}

func (repo *productRepository) GetByID(id int) (*model.Product, error) {
	var product model.Product
	query := `
		SELECT id, name, price, created_at
		FROM products
		WHERE id = ? AND archived_at IS NULL
	`
	err := repo.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (repo *productRepository) Update(product *model.Product) error {
	query := `
		UPDATE products 
		SET name = ?, price = ?
		WHERE id = ? AND archived_at IS NULL
	`
	result, err := repo.db.Exec(query, product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (repo *productRepository) Archive(id int) error {
	query := `
		UPDATE products
		SET archived_at = CURRENT_TIMESTAMP
		WHERE id = ? AND archived_at IS NULL
	`
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
