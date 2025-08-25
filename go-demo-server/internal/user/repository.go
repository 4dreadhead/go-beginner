package user

import (
	"goadvancedserver/pkg/db"
)

type Repository struct {
	database *db.DB
}

func NewRepository(database *db.DB) *Repository {
	return &Repository{database: database}
}

func (repo *Repository) Create(user *User) (*User, error) {
	result := repo.database.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *Repository) FindByUsername(username string) (*User, error) {
	var user User
	result := repo.database.DB.First(&user, "email = ?", username)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
