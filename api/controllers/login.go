package controllers

import (
	"net/http"
  "../models"
	"../db"
	"github.com/labstack/echo"
  "golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
  dbc := c.Get("db").(*db.Client)
  newUser := models.User{}

  if bindErr := c.Bind(&newUser); bindErr != nil {
    return c.JSON(http.StatusInternalServerError, bindErr)
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

  return c.JSON(http.StatusOK, "OK")
}

func Login(c echo.Context) error {
  dbc := c.Get("db").(*db.Client)
  u := models.User{}

  if bindErr := c.Bind(&u); bindErr != nil {
    return c.JSON(http.StatusInternalServerError, bindErr)
  }


  cu, dbErr := dbc.GetUserByName(u.Uname)

  if dbErr != nil {
    return c.JSON(http.StatusInternalServerError, dbErr)
  }
  return c.JSON(http.StatusOK, cu)
}
