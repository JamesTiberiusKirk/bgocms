package db

import (
	"database/sql"
	"fmt"

	"../models"
)

func (c *Client) GetUserRows() ([]models.User, int, error) {
  dbc := c.dbc

	sqlStatement := "SELECT * FROM users;"
	rows, err := dbc.Query(sqlStatement)

  if err != nil {
		fmt.Println(err)
    return nil, 0, err
	}

	defer rows.Close()
  result := []models.User{}
  total := 0

  for rows.Next() {
    total += 1
		rowUser := models.User{}
		err2 := rows.Scan(&rowUser.ID, &rowUser.Uname, &rowUser.Pass)

    // Exit if we get an error
    if err2 != nil {
			return nil, 0, err2
		}

    result = append(result, rowUser)
	}

  return result, total, nil
}

func (c *Client) InsertUserRow(u models.User) (*sql.Rows, error){
  dbc := c.dbc

  sqlStatement := "INSERT INTO users (uname, pass) VALUES ($1, $2);"
  res, err := dbc.Query(sqlStatement, u.Uname, u.Pass)
  if err != nil {
    return nil, err
  } else {
    return res, nil
  }
}
