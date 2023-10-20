package app

import "github.com/kabbachokk/golang-test-3/model"

// MysqlRepository
type MysqlRepository interface {
	QueryOrderProductsByOrderId(ids []int) ([]*model.ProductOrder, error)
	QueryProductsById(ids []int) ([]*model.Product, error)
	QueryRacksById(ids []int) ([]*model.Rack, error)
	QueryProductRacksByProductId(ids []int) ([]*model.ProductRack, error)
}
