package db

import (
	"fmt"
	"log"
  "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Client struct {
	db        *sqlx.DB
  ex        sqlx.Ext
  committed bool
}

func NewClient(db *sqlx.DB) *Client {
  return &Client{
    db: db,
    ex: db,
  }
}

func Connect() (*Client, error){
	db, err := sqlx.Open("postgres", "user=postgres password=example dbname=bgocms sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
    c := NewCLient(&db)
    return c, nil
	}
	return nil, err
}
