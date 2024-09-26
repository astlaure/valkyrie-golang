package core

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func SetupDatabase() {
	DB = sqlx.MustConnect("mysql", os.Getenv("DATABASE_URI"))
}
