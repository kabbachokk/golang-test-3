package usecase

import (
	"ip.com/app"
)

// useCase
type useCase struct {
	repo app.MysqlRepository
}

// NewProductUC constructor
func NewUseCase(
	repo app.MysqlRepository,
) *useCase {
	return &useCase{repo}
}
