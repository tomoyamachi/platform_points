package resource

import (
	"net/http"
	"platform_points/model"

	"github.com/echo-contrib/sessions"
	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	token := c.FormValue("token")
	appName := c.FormValue("app_code")
	account := model.Authenticate(token, appName)
	if account.Id == 0 {
		response := map[string]interface{}{
			"statusCode": http.StatusMethodNotAllowed,
			"message":    "Not authorized user!!",
		}
		return c.JSON(http.StatusMethodNotAllowed, response)
	}

	session := sessions.Default(c)

	session.Set("account_id", account.Id)
	session.Save()
	return c.JSON(http.StatusOK, account)
}
