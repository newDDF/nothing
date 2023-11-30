package config

import "database/sql"

type Mysql struct {
}

func ConnectMysql() {
	dsn := "root:root@tcp(localhost:3306)testdb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
