package controllers

import(
  "io"
  "os"
  "net/http"
  "github.com/labstack/echo"
)

func UploadPicture(c echo.Context) error {
  return c.String(http.StatusOK, "Successfully uploaded")
}
