package app

import "github.com/kabbachokk/golang-test-3/model"

// MysqlRepository
type MysqlRepository interface {
	QueryOrderRacksByOrderID(ids []int) ([]*model.OrderRack, error)
}
