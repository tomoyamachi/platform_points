package route

import (
	"platform_accounts/db"
	myMw "platform_accounts/middleware"
	"platform_accounts/resource"

	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
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
