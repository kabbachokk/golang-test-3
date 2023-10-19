package usecase

import (
	"github.com/kabbachokk/golang-test-3/app"
	"github.com/kabbachokk/golang-test-3/model"
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

func (uc *useCase) GetOrderRacksByOrderID(ids []int) ([]*model.OrderRacks, error) {
	return uc.repo.QueryOrderRacksByOrderID(ids)
}
