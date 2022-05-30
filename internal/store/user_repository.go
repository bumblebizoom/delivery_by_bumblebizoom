package store

import "github.com/bumblebizoom/delivery_by_bumblebizoom/internal/models"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *models.User) (*models.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, passwordhash) VALUES (?, ?)",
		u.Email,
		u.PasswordHash,
	).Scan(&u.Id); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	return nil, nil
}
