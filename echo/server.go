package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	// Path param
	e.GET("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})
	// Query param
	e.GET("/show", func(c echo.Context) error {
		team := c.QueryParam("team")
		member := c.QueryParam("member")
		return c.String(http.StatusOK, fmt.Sprintf("team:%s, member:%s", team, member))
	})
	// Post form
	e.POST("/save", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		return c.String(http.StatusOK, fmt.Sprintf("name:%s, email:%s", name, email))
	})
	// Maulti part form data
	e.POST("/save2", func(c echo.Context) error {
		name := c.FormValue("name")
		avatar, err := c.FormFile("avatar")
		if err != nil {
			return err
		}
		src, err := avatar.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		dst, err := os.Create(avatar.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return err
		}
		return c.HTML(http.StatusOK, fmt.Sprintf("<b>Thank you! %s</b>", name))
	})
	// Post JSON
	e.POST("/users", func(c echo.Context) error {
		u := &User{}
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
	}, track)

	// serve static files
	e.Static("/static", "static")

	// Middlewares
	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":3000"))
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func track(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		println("request to /users")
		return next(c)
	}
}
