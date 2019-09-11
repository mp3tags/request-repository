package service

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"log"
	"os"
	"time"
)

func ConnectToDb() *sql.DB {
	var db *sql.DB
	var err error

	for {
		// "username:password@tcp(host:port)/db_name"
		conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_PORT"),
			os.Getenv("MYSQL_DB_NAME"))

		db, err = sql.Open("mysql", conn)

		if err != nil {
			log.Println(err)
			time.Sleep(10 * time.Second)
			continue
		}

		// migrations
		driver, err := mysql.WithInstance(db, &mysql.Config{})
		if err != nil {
			log.Println(err)
		}

		m, err := migrate.NewWithDatabaseInstance(
			"file://migrations",
			"mysql",
			driver,
		)
		if err != nil {
			log.Println(err)
		}

		err = m.Steps(2)
		if err != nil {
			log.Println(err)
		}

		break
	}

	return db
}
