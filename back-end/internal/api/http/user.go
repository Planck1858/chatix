package http

import (
	"encoding/json"
	"githab.com/Planck1858/chatix/back-end/internal/user"
	"githab.com/Planck1858/chatix/back-end/pkg/logging"
	"github.com/labstack/echo"
	"net/http"
)

const (
	UserApiPath = "/v1/user"

	GetAllUsersPath = "/"
	GetUserPath     = "/:id"
	UpdateUserPath  = "/"
	DeleteUserPath  = "/:id"

	logUserControllerPath = "user.controller."
)

type userController struct {
	log         logging.Logger
	userService user.Service
}

func NewUserController(log logging.Logger, service user.Service) *userController {
	return &userController{
		log:         log,
		userService: service,
	}
}

func (c *userController) Register(r *echo.Echo) {
	g := r.Group(UserApiPath)

	g.GET(GetAllUsersPath, c.getAllUsers)
	g.GET(GetUserPath, c.getUser)
	g.PUT(UpdateUserPath, c.updateUser)
	g.DELETE(DeleteUserPath, c.deleteUser)
}

/***** getAllUsers *****/
func (c *userController) getAllUsers(ec echo.Context) (err error) {
	funcName := "getAllUsers"
	defer func() {
		if err != nil {
			c.log.With(err).Error(logUserControllerPath + funcName + " finished with error")
		} else {
			c.log.Info(logUserControllerPath + funcName + " finished correctly")
		}
	}()
	c.log.Info(logUserControllerPath + funcName + " get request...")
	ctx := ec.Request().Context()

	users, err := c.userService.GetAllUsers(ctx)
	if err != nil {
		return err
	}

	return ec.JSON(http.StatusOK, users)
}

/***** getUser *****/
func (c *userController) getUser(ec echo.Context) (err error) {
	funcName := "getUser"
	defer func() {
		if err != nil {
			c.log.With(err).Error(logUserControllerPath + funcName + " finished with error")
		} else {
			c.log.Info(logUserControllerPath + funcName + " finished correctly")
		}
	}()
	c.log.Info(logUserControllerPath + funcName + " get request...")
	ctx := ec.Request().Context()

	id := ec.Param("id")

	u, err := c.userService.GetUser(ctx, id)
	if err != nil {
		return err
	}

	return ec.JSON(http.StatusOK, u)
}

/***** updateUser *****/
func (c *userController) updateUser(ec echo.Context) (err error) {
	funcName := "updateUser"
	defer func() {
		if err != nil {
			c.log.With(err).Error(logUserControllerPath + funcName + " finished with error")
		} else {
			c.log.Info(logUserControllerPath + funcName + " finished correctly")
		}
	}()
	c.log.Info(logUserControllerPath + funcName + " get request...")
	ctx := ec.Request().Context()

	dto := user.UpdateUserDTO{}
	err = json.NewDecoder(ec.Request().Body).Decode(&dto)

	err = c.userService.UpdateUser(ctx, &dto)
	if err != nil {
		return err
	}

	return ec.NoContent(http.StatusOK)
}

/***** deleteUser *****/
func (c *userController) deleteUser(ec echo.Context) (err error) {
	funcName := "deleteUser"
	defer func() {
		if err != nil {
			c.log.With(err).Error(logUserControllerPath + funcName + " finished with error")
		} else {
			c.log.Info(logUserControllerPath + funcName + " finished correctly")
		}
	}()
	c.log.Info(logUserControllerPath + funcName + " get request...")
	ctx := ec.Request().Context()

	id := ec.Param("id")

	err = c.userService.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return ec.NoContent(http.StatusOK)
}
