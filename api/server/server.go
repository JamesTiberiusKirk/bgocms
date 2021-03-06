package server

import (
  "bgocms/db"
	"bgocms/controllers"
  "bgocms/mw"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// ContextParams stores context parameters for server initialization
type ContextParams struct {
  DB *db.Client
}

func Init(dbc *db.Client) *echo.Echo {
	e := echo.New()
  contextParaps := ContextParams{DB: dbc}
	e.Use(
		createContext(contextParaps),
		middleware.Logger(),
		middleware.Gzip(),
		middleware.CORS(),
	)

  e = initRoutes(e)
  return e
}

func initRoutes(e *echo.Echo) *echo.Echo{

  g := e.Group("/admin")

  g.Use(middleware.BasicAuth(mw.Auth))

	g.POST("/articles", controllers.AddArticle)
	g.POST("/users", controllers.AddUser)

	e.GET("/test", controllers.Test)
	e.GET("/users", controllers.GetUsers)
  e.GET("/articles", controllers.GetArticles)
  e.GET("/article", controllers.GetActicleByID)

  e.POST("/register", controllers.Register)
  e.POST("/login", controllers.Login)

  return e
}

// ContextObjects attaches backend clients to the API context
func createContext(contextParams ContextParams) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", contextParams.DB)
			return next(c)
		}
	}
}
