package resource

import (
	"net/http"
	"platform_points/model"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/echo-contrib/sessions"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func GetTargetAccountPoints(c echo.Context) error {
	session := sessions.Default(c)
	accountIdSession := session.Get("account_id")

	// URLからaccountIdをintで取得
	accountIdUrl, _ := strconv.ParseInt(c.Param("account_id"), 0, 64)

	if accountIdSession != accountIdUrl {
		response := map[string]interface{}{
			"statusCode": http.StatusMethodNotAllowed,
			"message":    "Not authorized user!!",
		}
		return c.JSON(http.StatusMethodNotAllowed, response)
	}

	tx := c.Get("Tx").(*dbr.Tx)
	u_point := new(model.UPoints)
	if err := u_point.LoadTargetAccountPoints(tx, accountIdUrl); err != nil {
		logrus.Debug(err)
		return echo.NewHTTPError(fasthttp.StatusNotFound, "Target User Points does not exists.")
	}

	return c.JSON(http.StatusOK, u_point)
}
