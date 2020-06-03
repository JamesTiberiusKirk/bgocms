package models

import(
  "time"
)

type Article struct {
  ID          int       `json:"id" db:"id"`
  Author      int       `json:"author" db:"author"`
  Title       string    `json:"title" db:"title"`
  Body        string    `json:"body" db:"body"`
  Created     time.Time `json:"created" db:"created"`
  Last_edited time.Time `json:"last_edited" db:"last_edited"`
}

type ArticleResponce struct {
  Articles []Article  `json:"articles"`
  Total int           `json:"total"`
}

type ArticleAuthor struct {
  ID          int       `json:"id" db:"id"`
  AuthorId    int       `json:"author_id" db:"user_id"`
  AuthorUname string    `json:"author_uname" db:"uname"`
  Title       string    `json:"title" db:"title"`
  Body        string    `json:"body" db:"body"`
  Created     time.Time `json:"created" db:"created"`
  Last_edited time.Time `json:"last_edited" db:"last_edited"`

}
