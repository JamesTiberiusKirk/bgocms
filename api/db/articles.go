package db

import (
  "bgocms/models"
)

func (c *Client) GetArticleRows() ([]models.Article, int ,error) {
  dbc := c.db

  sql := `SELECT * FROM articles;`
  rows, queryErr := dbc.Query(sql)

  if queryErr != nil {
    return nil, 0, queryErr
  }

  results := []models.Article{}
  total := 0

  for rows.Next() {
    total += 1
    article := models.Article{}
    scanErr := rows.Scan(&article)
    if scanErr != nil {
      return nil, 0, scanErr
    }
    results = append(results, article)
  }

  return results, total, nil
}

func (c *Client) InsertArticleRow(a models.Article) error {
  dbc := c.db
  sql := `INSERT INTO articles (author, title, body, created, last_edited) VALUES ($1, $2, $3, $4, $5);`
  _, queryErr := dbc.Query(sql, a.Author, a.Title, a.Body, a.Created, a.Last_edited)

  if queryErr != nil {
    return queryErr
  }

  return nil
}


