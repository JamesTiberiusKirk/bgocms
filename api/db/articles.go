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
    scanErr := rows.Scan( &article.ID,
                          &article.Author,
                          &article.Title,
                          &article.Body,
                          &article.Created,
                          &article.Last_edited)
    if scanErr != nil {
      return nil, 0, scanErr
    }
    results = append(results, article)
  }

  return results, total, nil
}

func (c *Client) GetArticleRowByID(id string) (*models.Article,  error){
  dbc := c.db
  sql := `SELECT * FROM articles WHERE id=$1;`

  rows, queryErr := dbc.Query(sql, id)

  if queryErr != nil {
    return nil, queryErr
  }

  result := models.Article{}

  for rows.Next(){
    scanErr := rows.Scan( &result.ID,
      &result.Author,
      &result.Title,
      &result.Body,
      &result.Created,
      &result.Last_edited )
    if scanErr != nil {
      return nil, scanErr
    }
  }
  return &result, nil
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
