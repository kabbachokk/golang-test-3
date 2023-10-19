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

type Row struct {
	PrimaryRack    string
	OrderID        int
	ProductID      int
	ProductName    string
	Qty            int
	SecondaryRacks string
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

/*
SELECT product_id, GROUP_CONCAT(r.name SEPARATOR ' ')
FROM product_rack pr
INNER JOIN rack r ON r.id = pr.rack_id
GROUP BY product_id;
*/

/*
SELECT p.name, po.order_id, po.qty, pg.racks
FROM product p
INNER JOIN product_order po ON p.id = po.product_id
INNER JOIN (
    SELECT product_id, GROUP_CONCAT(r.name SEPARATOR ':') racks
		FROM product_rack pr
    	WHERE pr.primary != 1
		INNER JOIN rack r ON r.id = pr.rack_id
		GROUP BY product_id
) pg ON p.id = pg.product_id
INNER JOIN (
    SELECT product_id, GROUP_CONCAT(r.name SEPARATOR ':') racks
		FROM product_rack pr
		INNER JOIN rack r ON r.id = pr.rack_id
		GROUP BY product_id
) pg ON p.id = pg.product_id
*/

/*
CREATE TRIGGER after_product_rack_insert
AFTER INSERT
ON `product_rack` FOR EACH ROW
BEGIN
  	IF NEW.is_primary IS true THEN
    	INSERT INTO `product_rack_names` SET
    	product_id = NEW.product_id,
    	p_name = (SELECT name FROM rack WHERE id = NEW.rack_id LIMIT 1),
		s_name = NULL;
	ELSE
		UPDATE `product_rack_names` SET
		s_name = CONCAT_WS(',', s_name, (SELECT name FROM rack WHERE id = NEW.rack_id LIMIT 1))
		WHERE product_id = NEW.product_id;
	END IF;
END;
*/
