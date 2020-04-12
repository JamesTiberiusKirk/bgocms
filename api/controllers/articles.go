package controllers

import(
  "time"
  "net/http"
  "bgocms/models"
  "bgocms/db"
  "github.com/labstack/echo"
)

func GetArticles(c echo.Context) error {
  dbc := c.Get("db").(*db.Client)
  articles, total, dbErr := dbc.GetArticleRows()
  if dbErr != nil {
    return c.JSON(http.StatusInternalServerError, dbErr.Error())
  }
  return c.JSON(http.StatusOK, models.ArticleResponce{
    Articles: articles,
    Total: total,
  })
}

func AddArticle(c echo.Context) error {
  dbc := c.Get("db").(*db.Client)
  newArticle := models.Article{}

  bindErr := c.Bind(&newArticle)
  if bindErr != nil {
    return c.JSON(http.StatusBadRequest, bindErr.Error())
  }

  timestamp := time.Now()
  newArticle.Created = timestamp
  newArticle.Last_edited = timestamp

  dbErr := dbc.InsertArticleRow(newArticle)
  if dbErr != nil {
    return c.JSON(http.StatusInternalServerError, dbErr.Error())
  }
 return c.String(http.StatusOK, "Added")
}
