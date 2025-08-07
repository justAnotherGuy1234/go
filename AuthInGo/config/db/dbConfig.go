package config

import (
	env "AuthInGo/config/env"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func SetupDb() (*sql.DB, error) {
	cfg := mysql.NewConfig()

	cfg.User = env.GetString("DB_USER", "root")
	cfg.Net = "tcp"
	cfg.Addr = env.GetString("DB_ADDR", "127.0.0.1:3306")
	cfg.Passwd = env.GetString("DB_PASSWORD", "")
	cfg.DBName = env.GetString("DE_NAME", "goTest")

	db, err := sql.Open("mysql", cfg.FormatDSN())
	fmt.Println("cfg format", cfg.FormatDSN())

	if err != nil {
		fmt.Println("error connecting to db ", err)
		return nil, err
	}

	pingErr := db.Ping()

	if pingErr != nil {
		fmt.Println("error pinging database", pingErr)
		return nil, pingErr
	}

	fmt.Println("connected to db", cfg.DBName)

	return db, nil
}
