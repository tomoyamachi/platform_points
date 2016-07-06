package resource

import (
	"net/http"
	"platform_points/model"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	token := c.FormValue("token")
	appName := c.FormValue("app_code")
	account := model.Authenticate(token, appName)
	return c.JSON(http.StatusOK, account)
}
