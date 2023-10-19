package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/kabbachokk/golang-test-3/model"
)

// mysqlRepo
type mysqlRepo struct {
	conn *sql.DB
}

// NewMysqlRepo
func NewMysqlRepo(conn *sql.DB) *mysqlRepo {
	return &mysqlRepo{conn}
}

func (r *mysqlRepo) QueryOrderRacksByOrderID(ids []int) ([]*model.OrderRacks, error) {
	//ids := []int{10, 11, 14, 15}
	sql := `
	SELECT prn.p_name, po.order_id, p.id, p.name, po.qty, prn.s_name
		FROM product_order po
		LEFT JOIN product p ON p.id = po.product_id
		LEFT JOIN product_rack_names prn ON prn.product_id = po.product_id
		WHERE po.order_id IN (%s)
	ORDER BY prn.p_name, po.order_id
	`
	sql = fmt.Sprintf(sql,
		strings.Join(strings.Split(strings.Repeat("?", len(ids)), ""), ","))

	stmt, err := r.conn.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	parm := make([]interface{}, len(ids))
	for i := range ids {
		parm[i] = ids[i]
	}
	rows, err := stmt.Query(parm...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*model.OrderRacks
	for rows.Next() {
		var m model.OrderRacks
		rows.Scan(&m.PrimaryRack, &m.OrderID, &m.ProductID, &m.ProductName, &m.Qty, &m.SecondaryRacks)
		res = append(res, &m)
	}

	return res, nil
}
