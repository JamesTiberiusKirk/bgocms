package main

import (
  "fmt"
	"bgocms/db"
	"bgocms/server"
	_ "github.com/lib/pq"
)

func main() {
	dbc, err := db.Connect()

  if err != nil {
    fmt.Println("error")
    return
  }

	s := server.Init(dbc)

	s.Start(":8000")
}
