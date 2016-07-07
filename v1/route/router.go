package route

import (
	"platform_points/db"
	"platform_points/handler"
	myMw "platform_points/middleware"
	"platform_points/v1/resource"

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

	//store, err := sessions.NewRedisStore(32, "tcp", "pointsession.iyjz8a.0001.apne1.cache.amazonaws.com:6379", "", []byte("secret"))
	//logrus.Debug(store.Get())
	e.Use(sessions.Sessions("PHPSESSID", store))

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
