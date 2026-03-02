package user

import (
	"database/sql"
	"jwt-auth/internal/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	// We can use ORM like GORM or sqlx
	// For now, we will use raw SQL
	user := new(types.User)
	err := s.db.QueryRow("SELECT id, email, password, first_name, last_name, created_at, updated_at, is_active FROM users WHERE email = ?", email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsActive,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	user := new(types.User)
	err := s.db.QueryRow("SELECT id, email, password, first_name, last_name, created_at, updated_at, is_active FROM users WHERE id = ?", id).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsActive,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (email, password, first_name, last_name) VALUES (?, ?, ?, ?)",
		user.Email, user.Password, user.FirstName, user.LastName)
	if err != nil {
		return err
	}

	return nil
}
