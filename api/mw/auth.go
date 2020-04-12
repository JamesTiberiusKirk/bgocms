package mw

import (
  "bgocms/db"
  "github.com/labstack/echo"
  "golang.org/x/crypto/bcrypt"
)

func Auth(username,password string, c echo.Context) (bool, error){
  dbc := c.Get("db").(*db.Client)

  cu, dbErr := dbc.GetUserByName(username)
  if dbErr != nil {
    return false, dbErr
  }
  compareErr := bcrypt.CompareHashAndPassword([]byte(cu.Pass), []byte(password))
  if compareErr != nil {
    return false, nil
  }

  return true, nil
}
