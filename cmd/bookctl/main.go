package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/olivere/bookstore/book"
)

const (
	DefaultMySQL_URL = "go:go@tcp(localhost:3306)/bookstore?loc=UTC&parseTime=true&multiStatements=true"
)

func main() {
	flag.Parse()

	// Get data source name to connect to MySQL
	dsn := os.Getenv("MYSQL_URL")
	if dsn == "" {
		dsn = DefaultMySQL_URL
	}

	// Connect to MySQL
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database server
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	log.Print("connected...")

	bookRepo := book.NewMySQLRepository(db)

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	switch flag.Arg(0) {
	case "read":
		id, err := strconv.ParseInt(flag.Arg(1), 10, 64)
		if err != nil {
			log.Fatalf("read needs a numerical id")
		}
		if err := readBook(ctx, bookRepo, id); err != nil {
			log.Fatal(err)
		}
	default:
		flag.Usage()
	}
}
