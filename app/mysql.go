package app

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connect_db(telegramId string) string {

	db, err := sql.Open("mysql", ":@tcp(192.168.0.100:3306)/totp")

	if err != nil {
		panic(err)
	}

	defer db.Close()
	rows, err := db.Query("select security_key from users where telegram = ?", telegramId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var Key string

	for rows.Next() {

		err := rows.Scan(&Key)
		if err != nil {
			fmt.Println(err)
			continue
		}

	}
	return Key
}

