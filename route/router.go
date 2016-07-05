package route

import (
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
	"github.com/tomoyamachi/platform_accounts/db"
	myMw "github.com/tomoyamachi/platform_accounts/middleware"
	"github.com/tomoyamachi/platform_accounts/resource"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Debug()
	// Set MiddleWare
	e.Use(echoMw.Logger())
	e.Use(echoMw.Recover())
	e.Use(echoMw.Gzip())

	// Set Custom MiddleWare
	e.Use(myMw.TxMiddleware(db.Init()))

	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.Post("/members", resource.CreateMember)
		v1.Get("/members", resource.GetMembers)
		v1.Get("/members/:id", resource.GetMember)
	}
	return e
}
