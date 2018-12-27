package app

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	ctl "github.com/long-in/go-echo-sample/app/controllers"
	validator "gopkg.in/go-playground/validator.v9"
)

type Validator struct {
	validator *validator.Validate
}

var Server *echo.Echo

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

// init echo web server
func Init() {
	Server = echo.New()
	Server.Validator = &Validator{validator: validator.New()}

	// Middleware
	Server.Use(middleware.Logger())
	Server.Use(middleware.Recover())

	// Basic Auth All users
	//Server.Use(sv.BasicAuth())

	// CORS
	Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	//
	Server.POST("/v1/fullname", ctl.FullNameMaker)
	//Server.POST("/v1/fullname", ctl.FullNameMaker, sv.BasicAuth())

	//
	Server.GET("/v1/ok", func(c echo.Context) error {
		return c.String(200, "OK")
	})
}
