package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo/engine/standard"
	"github.com/tomoyamachi/platform_accounts/route"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	router := route.Init()
	router.Run(standard.New(":1323"))
}
