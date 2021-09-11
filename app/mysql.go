package app

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type MYSQL struct {
	DB *sql.DB
}

func (c *Config) NewMysqlConnection() *sql.DB  {
	db, err := sql.Open("mysql", c.MYSQL)
	if err != nil {
		panic(err)
	}
	return db
}

func (m *MYSQL) GetLogin(telegramId string) string {

	rows, err := m.DB.Query("select `user_login` from users where telegram = ?", telegramId)
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

