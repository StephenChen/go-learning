package main

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

// TestShouldUpdateStats sql 执行成功的测试用例
func TestShouldUpdateStats(t *testing.T) {
	// mock 一个 *sql.DB 对象, 不需要连接真实的数据库
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an err '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// mock 执行指定 SQL 语句时的返回结果
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO product_viewers").
		WithArgs(2, 3).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// 将 mock 的 DB 对象传入我们的函数中
	if err = recordStats(db, 2, 3); err != nil {
		t.Errorf("err was not expected while updating stats: %s", err)
	}

	// 确保期望的结果都满足
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfilfilled expactations: %s", err)
	}
}

// TestShouldRollbackStatUpdatesOnFailure sql 执行失败回滚的测试用例
func TestShouldRollbackStatUpdatesOnFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO product_viewers").
		WithArgs(2, 3).
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	// now we execute our method
	if err = recordStats(db, 2, 3); err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	// we make sure that all expectations were met
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
