package controllers

import (
	"fmt"
	"net/http"
  "../models"

	"../db"
	"github.com/labstack/echo"
)

// Test this is for testing packages
func Test(c echo.Context) error {
	fmt.Println("helloWorld")
	return c.String(http.StatusOK, "helloworld from Users")
}

//GetUsers gets all of the rows from users table
func GetUsers(c echo.Context) error {
	dbc := c.Get("db").(*db.Client)
	users, total, err := dbc.GetUserRows()

  if err != nil {
    return c.JSON(http.StatusInternalServerError, err)
  }

  return c.JSON(http.StatusOK, models.UsersResponce{
    Users:  users,
    Total:  total,
  })
}

//AddUser this is for posting a new user into users table
func AddUser(c echo.Context) error {
	dbc := c.Get("db").(*db.Client)
	u := models.User{}

	if err := c.Bind(&u); err != nil {
		return err
	}

  res, db_err := dbc.InsertUserRow(u)
  if db_err != nil {
    return c.JSON(http.StatusInternalServerError, db_err)
  }

  return c.JSON(http.StatusCreated, res)
}
