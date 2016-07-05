package main

import (
	"platform_accounts/route"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo/engine/standard"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	router := route.Init()
	router.Run(standard.New(":1323"))
}
