package route

import (
	"platform_points/db"
	"platform_points/handler"
	myMw "platform_points/middleware"
	"platform_points/v1/resource"

	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Debug()

	// Set Bundle MiddleWare
	e.Use(echoMw.Logger())
	e.Use(echoMw.Gzip())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
	e.SetHTTPErrorHandler(handler.JSONHTTPErrorHandler)

	// Set Custom MiddleWare
	e.Use(myMw.TransactionHandler(db.Init()))

	// Routes
	v1 := e.Group("/v1")
	{
		v1.POST("/members", resource.PostMember())
		v1.GET("/members", resource.GetMembers())
		v1.GET("/members/:id", resource.GetMember())
	}
	return e
}
