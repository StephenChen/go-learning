package main

import "database/sql"

// recordStats 记录用户浏览产品信息
func recordStats(db *sql.DB, userID, productID int64) error {
	// 开启事务
	// 操作 views 和 product_viewers 两张表
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	// 更新 product 表
	if _, err = tx.Exec("UPDATE products SET views = views + 1"); err != nil {
		return err
	}
	// product_viewers 表中插入一条数据
	if _, err = tx.Exec("INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?)", userID, productID); err != nil {
		return err
	}
	return nil
}

func main() {
	// 注意：测试的过程中并不需要真正的连接
	db, err := sql.Open("mysql", "root@/blog")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// userID 为 1 的用户浏览了 productID 为 5 的产品
	if err = recordStats(db, 1, 5); err != nil {
		panic(err)
	}
}
