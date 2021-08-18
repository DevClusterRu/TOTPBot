package app

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

func Connect_db(telegramId string) string {

	db, err := sql.Open("mysql", "user:123456@tcp(10.175.255.30)/new_hdesk")

	if err != nil {
		panic(err)
	}

	defer db.Close()
	rows, err := db.Query("select `user_login` from users") // where telegram = ?", telegramId
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var key string

	for rows.Next() {

		err := rows.Scan(&key)
		if err != nil {
			fmt.Println(err)
			continue
		}

	}

	if strings.Contains(key,`\`){
		key = key[strings.Index(key, `\`)+1:]
	}

	return key
}

