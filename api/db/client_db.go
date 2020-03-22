package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Client struct {
	dbc *sql.DB
}

func NewClient() (*Client, error){
	db, err := sql.Open("postgres", "user=postgres password=example dbname=bgocms sslmode=disable")
  c := &Client{dbc:db}
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
    c.dbc = db
		return c, nil
	}
	return nil, err
}
