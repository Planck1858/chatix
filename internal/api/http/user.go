package http

import (
	"encoding/json"
	"githab.com/Planck1858/chatix/internal/user"
	"githab.com/Planck1858/chatix/pkg/logging"
	"github.com/labstack/echo"
	"net/http"
)

const (
	UserApiPath = "/v1/user"

	GetAllUsersPath = "/"
	GetUserPath     = "/:id"
	CreateUserPath  = "/"
	UpdateUserPath  = "/"
	DeleteUserPath  = "/:id"
)

type userController struct {
	log     logging.Logger
	service user.Service
}

func NewUserController(log logging.Logger, service user.Service) *userController {
	return &userController{
		log:     log,
		service: service,
	}
}

func (c *userController) Register(r *echo.Echo) {
	g := r.Group(UserApiPath)

	g.GET(GetAllUsersPath, c.getAllUsers)
	g.GET(GetUserPath, c.getUser)
	g.POST(CreateUserPath, c.createUser)
	g.PUT(UpdateUserPath, c.updateUser)
	g.DELETE(DeleteUserPath, c.deleteUser)
}

/***** getAllUsers *****/
func (c *userController) getAllUsers(ec echo.Context) (err error) {
	funcName := "getAllUsers"
	defer func() {
		if err != nil {
			c.log.With(err).Error("user.controller." + funcName + " finished with error")
		} else {
			c.log.Info("user.controller." + funcName + " finished correctly")
		}
	}()
	c.log.Info("user.controller." + funcName + " get request...")
	ctx := ec.Request().Context()

	users, err := c.service.GetAllUsers(ctx)
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
			c.log.With(err).Error("user.controller." + funcName + " finished with error")
		} else {
			c.log.Info("user.controller." + funcName + " finished correctly")
		}
	}()
	c.log.Info("user.controller." + funcName + " get request...")
	ctx := ec.Request().Context()

	id := ec.Param("id")

	u, err := c.service.GetUser(ctx, id)
	if err != nil {
		return err
	}

	return ec.JSON(http.StatusOK, u)
}

type createUserResponse struct {
	Id string `json:"id"`
}

/***** createUser *****/
func (c *userController) createUser(ec echo.Context) (err error) {
	funcName := "createUser"
	defer func() {
		if err != nil {
			c.log.With(err).Error("user.controller." + funcName + " finished with error")
		} else {
			c.log.Info("user.controller." + funcName + " finished correctly")
		}
	}()
	c.log.Info("user.controller." + funcName + " get request...")
	ctx := ec.Request().Context()

	dto := user.CreateUserDTO{}
	err = json.NewDecoder(ec.Request().Body).Decode(&dto)

	id, err := c.service.CreateUser(ctx, &dto)
	if err != nil {
		return err
	}

	return ec.JSON(http.StatusCreated, createUserResponse{Id: id})
}

/***** updateUser *****/
func (c *userController) updateUser(ec echo.Context) (err error) {
	funcName := "updateUser"
	defer func() {
		if err != nil {
			c.log.With(err).Error("user.controller." + funcName + " finished with error")
		} else {
			c.log.Info("user.controller." + funcName + " finished correctly")
		}
	}()
	c.log.Info("user.controller." + funcName + " get request...")
	ctx := ec.Request().Context()

	dto := user.UpdateUserDTO{}
	err = json.NewDecoder(ec.Request().Body).Decode(&dto)

	err = c.service.UpdateUser(ctx, &dto)
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
			c.log.With(err).Error("user.controller." + funcName + " finished with error")
		} else {
			c.log.Info("user.controller." + funcName + " finished correctly")
		}
	}()
	c.log.Info("user.controller." + funcName + " get request...")
	ctx := ec.Request().Context()

	id := ec.Param("id")

	err = c.service.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return ec.NoContent(http.StatusOK)
}
