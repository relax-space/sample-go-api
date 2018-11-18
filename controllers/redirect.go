package controllers

import (
	"fmt"
	"net/http"

	"github.com/pangpanglabs/goutils/httpreq"

	"github.com/labstack/echo"
)

type RedirectController struct {
}

func (d RedirectController) Init(g *echo.Group) {
	g.GET("/real", d.Real)
	g.GET("", d.Get)
	g.POST("", d.Post)
	g.GET("/1", d.GetGet)
	g.POST("/1", d.PostGet)
}

func (d RedirectController) GetGet(c echo.Context) error {
	var v struct {
		Name string
	}
	status, err := httpreq.New(http.MethodGet, "http://localhost:8080/jump", nil).Call(&v)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	fmt.Println(status, err)
	return c.JSON(http.StatusOK, v)
}

func (d RedirectController) PostGet(c echo.Context) error {
	var v struct {
		Name string
	}
	status, err := httpreq.New(http.MethodPost, "http://localhost:8080/jump", nil).Call(&v)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	fmt.Println(status, err)
	return c.JSON(http.StatusOK, v)
}

func (d RedirectController) Real(c echo.Context) error {
	return c.JSON(http.StatusOK, struct {
		Name string
	}{
		"xiao",
	})
}

// http://localhost:8080/jump
func (d RedirectController) Get(c echo.Context) error {
	return c.Redirect(http.StatusFound, "http://localhost:8080/jump/real")
}

func (d RedirectController) Post(c echo.Context) error {
	return c.Redirect(http.StatusFound, "http://localhost:8080/jump/real")
}
