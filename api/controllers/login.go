package controllers

import (
  "fmt"
	"net/http"
  "bgocms/models"
	"bgocms/db"
	"github.com/labstack/echo"
  "golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
  dbc := c.Get("db").(*db.Client)
  newUser := models.User{}
  fmt.Printf("REGISTER")

  if bindErr := c.Bind(&newUser); bindErr != nil {
    return c.JSON(http.StatusBadRequest, bindErr)
  }

  hash, hashErr := bcrypt.GenerateFromPassword([]byte(newUser.Pass), 10)

  if hashErr != nil {
    return c.JSON(http.StatusInternalServerError, hashErr)
  }

  newUser.Pass = string(hash)
  dbErr := dbc.InsertUserRow(newUser)

  if dbErr != nil {
    return c.JSON(http.StatusInternalServerError, dbErr)
  }

  return c.JSON(http.StatusCreated, "OK")
}

func Login(c echo.Context) error {
  dbc := c.Get("db").(*db.Client)
  u := models.User{}

  if bindErr := c.Bind(&u); bindErr != nil {
    return c.JSON(http.StatusBadRequest, bindErr)
  }


  cu, dbErr := dbc.GetUserByName(u.Uname)

  if dbErr != nil {
    return c.JSON(http.StatusInternalServerError, dbErr)
  }

  compareErr := bcrypt.CompareHashAndPassword([]byte(cu.Pass), []byte(u.Pass))
  if compareErr != nil {
    return c.JSON(http.StatusUnauthorized, compareErr.Error())
  }

  return c.JSON(http.StatusOK, "Logged In")
}
