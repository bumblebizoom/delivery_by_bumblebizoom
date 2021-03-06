package repositories

import "GoProject_1/repositories/models"

type UserRepositories interface {
	GetByEmail(email string) *models.User
	Create(user *models.User) (*models.User, error)
}

type SupplierRepositories interface {
	GetAll() ([]*models.Supplier, error)
}

type ProductsRepositories interface {
	GetByCategory(category string) ([]*models.Product, error)
}

type OrdersRepositories interface {
	GetByID(ID uint64) (*models.Order, error)
}
