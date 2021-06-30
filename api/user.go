package api

import (
	"github.com/Planck1858/chatix/pkg/user"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
)

type userController struct {
	service user.Service
}

func NewUserController(service user.Service) *userController {
	return &userController{service: service}
}

func (c *userController) getUser(ec echo.Context) error {
	user := c.service.GetUser()
	if user == nil {
		return errors.New("no user")
	}

	return ec.JSON(http.StatusOK, user)
}

func (c *userController) Mount(r *echo.Echo) {
	r.Group("/user").
		Add(http.MethodGet, "/", c.getUser)
}
