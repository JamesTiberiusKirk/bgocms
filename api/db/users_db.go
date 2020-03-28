package db

import (
//  "database/sql"
  "bgocms/models"
)

func (c *Client) GetUserRows() ([]models.User, int, error) {
  dbc := c.db

	sqlStatement := "SELECT * FROM users;"
	rows, err := dbc.Query(sqlStatement)

  if err != nil {
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

func (c *Client) InsertUserRow(u models.User) error {
  dbc := c.db

  sqlStatement := `INSERT INTO users (uname, pass) VALUES ($1, $2);`
  //tx, err := dbc.Begin()

  _, err := dbc.Exec(sqlStatement, u.Uname, u.Pass)
  if err != nil {
    panic(err)
  } else {
    return nil
  }
}

func (c *Client) GetUserByName(searchName string) (*models.User, error){
  dbc := c.db
  sqlStatement := `SELECT * FROM users WHERE users.uname=$1;`
  u := models.User{}

  err := dbc.QueryRow(sqlStatement, searchName).Scan(&u.ID, &u.Uname, &u.Pass)

  if err != nil {
    panic(err)
  }

  return &u, nil
}
