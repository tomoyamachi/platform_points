package resource

import (
	"strconv"

	"platform_points/model"

	"github.com/Sirupsen/logrus"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func GetMPoint() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		number, _ := strconv.ParseInt(c.Param("id"), 0, 64)

		tx := c.Get("Tx").(*dbr.Tx)

		m_point := new(model.MPoint)
		if err := m_point.Load(tx, number); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusNotFound, "MPoints does not exists.")
		}
		return c.JSON(fasthttp.StatusOK, m_point)
	}
}

func GetMPoints() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		tx := c.Get("Tx").(*dbr.Tx)
		m_points := new(model.MPoints)
		if err = m_points.Load(tx); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusNotFound, "MPoint does not exists.")
		}

		return c.JSON(fasthttp.StatusOK, m_points)
	}
}
