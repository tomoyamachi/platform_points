package main

import (
	"github.com/labstack/echo/engine/standard"
	"github.com/tomoyamachi/platform_accounts/route"
)

func main() {
	router := route.Init()
	router.Run(standard.New(":1323"))
}
