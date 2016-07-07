package resource

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/echo-contrib/sessions"
	"github.com/labstack/echo"
)

func GetTargetAccountPoints(c echo.Context) error {
	session := sessions.Default(c)
	accountIdSession := session.Get("account_id")
	accountIdUrl := c.Param("account_id")

	logrus.Debug(accountIdSession)
	logrus.Debug(accountIdUrl)

	if accountIdSession != accountIdUrl {
		response := map[string]interface{}{
			"statusCode": http.StatusMethodNotAllowed,
			"message":    "Not authorized user!!",
		}
		return c.JSON(http.StatusMethodNotAllowed, response)
	}

	return c.JSON(http.StatusOK, accountIdSession)
}
