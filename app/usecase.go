package app

import "github.com/kabbachokk/golang-test-3/model"

// UseCase
type UseCase interface {
	GetOrderRacksByOrderID([]int) ([]*model.OrderRacks, error)
}
