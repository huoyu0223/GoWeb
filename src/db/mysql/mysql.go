package mysqlDB

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbs *sql.DB
)

func NewMysqlConn(username string, password string, ip string, port int16, DBName string, conPool int) error {
	dbs, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, ip, port, DBName))
	if err != nil {
		return err
	}
	if conPool < 10 {
		conPool = 10
	}
	if conPool > 100 {
		conPool = 100
	}
	dbs.SetMaxOpenConns(conPool)
	dbs.SetConnMaxIdleTime(10)
	dbs.SetConnMaxLifetime(time.Hour)
	return err
}

func Close() {
	dbs.Close()
}

func Query(sql string, args ...interface{}) (*sql.Rows, error) {
	rows, err := dbs.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func Exec(sql string, args ...interface{}) (sql.Result, error) {
	result, err := dbs.Exec(sql, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func test() {
	err := NewMysqlConn("username", "password", "host", 3389, "database", 20)
	if err != nil {
		panic(err.Error())
	}
	defer Close()

	// 查询数据库
	rows, err := Query("SELECT * FROM table WHERE id=?", 1)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		// ...
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(id, name)
	}

	// 插入数据
	result, err := Exec("INSERT INTO table(name, age) VALUES(?, ?)", "Bob", 20)
	if err != nil {
		panic(err.Error())
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("LastInsertId:", lastInsertID, "RowsAffected:", affectedRows)
}
