package controllers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type (
	Name struct {
		Fname string `json:"fname" validate:"required"`
		Lname string `json:"lname" validate:"required"`
	}

	FullName struct {
		FullName string `json:"fullname"`
	}
)

func FullNameMaker(c echo.Context) error {
	r := new(Name)
	if err := c.Bind(r); err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
	}

	if err := c.Validate(r); err != nil {
		return c.String(400, fmt.Sprintf("%s", err))
	}

	fn := []FullName{
		{FullName: r.Lname + r.Fname},
	}

	return c.JSON(http.StatusOK, fn)
}
