package route

import (
	"platform_points/apps/v1/resource"
	"platform_points/db"
	"platform_points/handler"
	myMw "platform_points/middleware"

	"github.com/echo-contrib/sessions"
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
	//store := sessions.NewCookieStore([]byte("123456"))
	store, err := sessions.NewRedisStore(32, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		panic(err)
	}

	e.Use(sessions.Sessions("pointsession", store))

	// Routes
	v1 := e.Group("/v1")
	{
		v1.GET("/m_points", resource.GetMPoints())
		v1.GET("/m_points/:id", resource.GetMPoint())
		v1.POST("/login", resource.Login)

		// need sessions
		v1.GET("/accounts/:account_id/points", resource.GetTargetAccountPoints)
	}
	return e
}
