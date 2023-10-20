package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/kabbachokk/golang-test-3/model"
	"github.com/kabbachokk/golang-test-3/utils"
)

// mysqlRepo
type mysqlRepo struct {
	conn *sql.DB
}

// NewMysqlRepo
func NewMysqlRepo(conn *sql.DB) *mysqlRepo {
	return &mysqlRepo{conn}
}

func (r *mysqlRepo) QueryOrderProductsByOrderId(ids []int) ([]*model.ProductOrder, error) {
	sql := `
	SELECT product_id, order_id, qty
	FROM product_order
	WHERE order_id IN (%s)
	`
	sql = fmt.Sprintf(sql,
		strings.Join(strings.Split(strings.Repeat("?", len(ids)), ""), ","))

	parm := utils.Int2interface(ids)
	rows, err := r.conn.Query(sql, parm...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*model.ProductOrder
	for rows.Next() {
		var m model.ProductOrder
		rows.Scan(&m.ProductId, &m.OrderId, &m.Qty)
		res = append(res, &m)
	}

	return res, nil
}

func (r *mysqlRepo) QueryProductsById(ids []int) ([]*model.Product, error) {
	sql := `
	SELECT id, name
	FROM product
	WHERE id IN (%s)
	`
	sql = fmt.Sprintf(sql,
		strings.Join(strings.Split(strings.Repeat("?", len(ids)), ""), ","))

	parm := utils.Int2interface(ids)
	rows, err := r.conn.Query(sql, parm...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*model.Product
	for rows.Next() {
		var m model.Product
		rows.Scan(&m.ID, &m.Name)
		res = append(res, &m)
	}

	return res, nil
}

func (r *mysqlRepo) QueryRacksById(ids []int) ([]*model.Rack, error) {
	sql := `
	SELECT id, name
	FROM rack
	WHERE id IN (%s)
	ORDER BY name
	`
	sql = fmt.Sprintf(sql,
		strings.Join(strings.Split(strings.Repeat("?", len(ids)), ""), ","))

	parm := utils.Int2interface(ids)
	rows, err := r.conn.Query(sql, parm...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*model.Rack
	for rows.Next() {
		var m model.Rack
		rows.Scan(&m.ID, &m.Name)
		res = append(res, &m)
	}

	return res, nil
}

func (r *mysqlRepo) QueryProductRacksByProductId(ids []int) ([]*model.ProductRack, error) {
	sql := `
	SELECT product_id, rack_id, is_primary
	FROM product_rack
	WHERE product_id IN (%s)
	`
	sql = fmt.Sprintf(sql,
		strings.Join(strings.Split(strings.Repeat("?", len(ids)), ""), ","))

	parm := utils.Int2interface(ids)
	rows, err := r.conn.Query(sql, parm...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*model.ProductRack
	for rows.Next() {
		var m model.ProductRack
		rows.Scan(&m.ProductId, &m.RackId, &m.IsPrimary)
		res = append(res, &m)
	}

	return res, nil
}
