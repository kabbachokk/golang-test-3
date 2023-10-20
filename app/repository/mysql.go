package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/kabbachokk/golang-test-3/model"
	"github.com/samber/lo"
)

// mysqlRepo
type mysqlRepo struct {
	conn *sql.DB
}

// NewMysqlRepo
func NewMysqlRepo(conn *sql.DB) *mysqlRepo {
	return &mysqlRepo{conn}
}

func (r *mysqlRepo) QueryOrderRacksByOrderID(ids []int) ([]*model.OrderRack, error) {
	//ids := []int{10, 11, 14, 15}
	sql := `
	SELECT r.name, po.order_id, po.product_id, p.name, po.qty 
	FROM 
		product_order AS po, 
		product_rack AS pr, 
		rack AS r,
		product AS p
	WHERE 
		pr.rack_id = r.id AND 
		pr.product_id = po.product_id AND 
		po.product_id = p.id AND 
		pr.is_primary IS true AND 
		po.order_id IN (%s)
	ORDER BY r.name;
	`
	sql = fmt.Sprintf(sql,
		strings.Join(strings.Split(strings.Repeat("?", len(ids)), ""), ","))

	parm := make([]interface{}, len(ids))
	for i := range ids {
		parm[i] = ids[i]
	}
	rows, err := r.conn.Query(sql, parm...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productIds []int

	var res []*model.OrderRack
	for rows.Next() {
		var m model.OrderRack
		rows.Scan(&m.PrimaryRack, &m.OrderID, &m.ProductID, &m.ProductName, &m.Qty)
		productIds = append(productIds, m.ProductID)
		res = append(res, &m)
	}

	productIds = lo.Uniq(productIds)
	secondaryParm := make([]interface{}, len(productIds))
	for i := range productIds {
		secondaryParm[i] = productIds[i]
	}

	secondarySql := `
	SELECT pr.product_id, GROUP_CONCAT(r.name) 
	FROM product_rack AS pr, rack AS r 
	WHERE pr.rack_id = r.id AND pr.is_primary IS NULL AND pr.product_id IN (%s)
	GROUP BY pr.product_id;
	`
	secondarySql = fmt.Sprintf(secondarySql,
		strings.Join(strings.Split(strings.Repeat("?", len(productIds)), ""), ","))

	secondaryRows, err := r.conn.Query(secondarySql, secondaryParm...)
	if err != nil {
		return nil, err
	}
	defer secondaryRows.Close()

	secondaryMap := map[int]string{}
	for secondaryRows.Next() {
		r := struct {
			ProductID int
			Names     string
		}{}
		secondaryRows.Scan(&r.ProductID, &r.Names)
		secondaryMap[r.ProductID] = r.Names
	}

	for _, v := range res {
		v.SecondaryRacks = secondaryMap[v.ProductID]
	}

	return res, nil
}
