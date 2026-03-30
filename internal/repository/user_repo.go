package repository

import (
	"database/sql"
	"ims/internal/model"
)

type UserRepository interface {
	Create(user *model.User) error
	GetByID(id int) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Update(user *model.User) error
	Archive(id int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) Create(user *model.User) error {
	query := `
		INSERT INTO users (name, email)
		VALUES (?, ?)
		RETURNING id, created_at
	`
	err := repo.db.QueryRow(query, user.Name, user.Email).Scan(&user.ID, &user.CreatedAt)
	return err
}

func (repo *userRepository) GetByID(id int) (*model.User, error) {
	var user model.User
	query := `
		SELECT id, name, email, created_at
		FROM users
		WHERE id = ? AND archived_at IS NULL
	`
	err := repo.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	query := `
		SELECT id, name, email, created_at
		FROM users
		WHERE email = ? AND archived_at IS NULL
	`
	err := repo.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) Update(user *model.User) error {
	query := `
		UPDATE users 
		SET name = ? 
		SET email = ?
		WHERE id = ? AND archived_at IS NULL
	`
	result, err := repo.db.Exec(query, user.Name, user.Email, user.ID)
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

func (repo *userRepository) Archive(id int) error {
	query := `
		UPDATE users
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
