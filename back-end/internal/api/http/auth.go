package http

import (
	"encoding/json"
	"githab.com/Planck1858/chatix/back-end/internal/auth"
	"githab.com/Planck1858/chatix/back-end/pkg/logging"
	"github.com/labstack/echo"
	"net/http"
)

const (
	AuthApiPath = "/v1/auth"

	SignUpPath = "/sign-up"
	SignInPath = "/sign-in"

	logAuthControllerPath = "auth.controller."
)

type authController struct {
	log         logging.Logger
	authService auth.Service
}

func NewAuthController(log logging.Logger, authService auth.Service) *authController {
	return &authController{
		log:         log,
		authService: authService,
	}
}

func (c *authController) Register(r *echo.Echo) {
	g := r.Group(AuthApiPath)

	g.POST(SignUpPath, c.signUp)
	g.POST(SignInPath, c.signIn)
}

/***** signUp *****/
type signUpResponse struct {
	userId string `json:"userId"`
}

func (c *authController) signUp(ec echo.Context) (err error) {
	funcName := "signUp"
	defer func() {
		if err != nil {
			c.log.With(err).Error(logAuthControllerPath + funcName + " finished with error")
		} else {
			c.log.Info(logAuthControllerPath + funcName + " finished correctly")
		}
	}()
	c.log.Info(logAuthControllerPath + funcName + " get request...")
	ctx := ec.Request().Context()

	dto := auth.SignUpDTO{}
	err = json.NewDecoder(ec.Request().Body).Decode(&dto)

	id, err := c.authService.SignUp(ctx, dto)
	if err != nil {
		return err
	}

	return ec.JSON(http.StatusOK, signUpResponse{userId: id})
}
