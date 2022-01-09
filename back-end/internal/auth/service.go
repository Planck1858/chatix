package auth

import (
	"context"
	"githab.com/Planck1858/chatix/back-end/internal/user"
	"githab.com/Planck1858/chatix/back-end/pkg/logging"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

const logServicePath = "auth.service."
const passwordHashCost = 14

type service struct {
	userService user.Service
	validator   *validator.Validate
}

func NewService(userService user.Service) Service {
	return &service{
		userService: userService,
		validator:   validator.New(),
	}
}

func (s *service) SignUp(ctx context.Context, dto *SignUpDTO) (userId string, err error) {
	funcName := "SignUp"
	log := logging.GetLogger()
	defer func() {
		if err != nil {
			log.With(err).Error(logServicePath + funcName + " finished with error")
		} else {
			log.Info(logServicePath + funcName + " finished correctly")
		}
	}()
	log.Info(logServicePath + funcName + " started...")

	err = s.validator.Struct(dto)
	if err != nil {
		return userId, errors.Wrap(err, "dto is invalid")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(dto.PasswordRaw), passwordHashCost)
	if err != nil {
		return userId, errors.Wrap(err, "password hash did not generate")
	}

	userId, err = s.userService.CreateUser(ctx, &user.CreateUserDTO{
		Login:    dto.Login,
		Name:     dto.Name,
		Email:    dto.Email,
		Role:     user.Role(dto.Role),
		Password: passwordHash,
	})
	if err != nil {
		return userId, errors.Wrap(err, "new user was not created")
	}

	return userId, nil
}

func (s *service) SignIn(ctx context.Context, dto *SignInDTO) (ok bool, err error) {
	funcName := "SignIn"
	log := logging.GetLogger()
	defer func() {
		if err != nil {
			log.With(err).Error(logServicePath + funcName + " finished with error")
		} else {
			log.Info(logServicePath + funcName + " finished correctly")
		}
	}()
	log.Info(logServicePath + funcName + " started...")

	err = s.validator.Struct(dto)
	if err != nil {
		return ok, errors.Wrap(err, "dto is invalid")
	}

	u, err := s.userService.GetUserByEmail(ctx, dto.Email)
	if err != nil {
		return ok, errors.Wrapf(err, "user with email '%v' was not found", dto.Email)
	}

	if u == nil || !u.DeletedAt.IsZero() {
		return false, errors.New("user is empty")
	}

	err = bcrypt.CompareHashAndPassword(u.Password, []byte(dto.Password))
	if err != nil {
		return false, errors.New("password in incorrect")
	}

	return true, nil
}
