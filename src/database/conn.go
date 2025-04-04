package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)
func ConnetionMysql() *sql.DB {
    host := os.Getenv("DATABASE_HOST")
    pwd := os.Getenv("DATABASE_PASSWORD")
    port := os.Getenv("DATABASE_PORT")
    dbname := os.Getenv("DATABASE_NAME")
    user := os.Getenv("DATABASE_USER")
    config := mysql.Config{
        User: user,
        Passwd: pwd,
        Net:  "tcp",
        Addr: fmt.Sprintf("%s:%s", host, port),
        DBName: dbname,
    }
    db, err := sql.Open("mysql", config.FormatDSN())
    if err != nil {
        panic(err)
    }
    if err := db.Ping(); err != nil {
        panic(err)
    }
    return db
}
