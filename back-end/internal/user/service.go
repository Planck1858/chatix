package user

import (
	"context"
	"githab.com/Planck1858/chatix/back-end/internal/storage/user"
	"githab.com/Planck1858/chatix/back-end/pkg/logging"
	"github.com/go-playground/validator/v10"
	//validation "github.com/go-ozzo/ozzo-validation/v4"
)

const logServicePath = "user.service."

type service struct {
	rep       user.Repository
	validator *validator.Validate
}

func NewUserService(ur user.Repository) Service {
	return Service(&service{
		rep:       ur,
		validator: validator.New(),
	})
}

func (s *service) GetAllUsers(ctx context.Context) (u []*User, err error) {
	funcName := "GetAllUsers"
	log := logging.GetLogger()
	defer func() {
		if err != nil {
			log.With(err).Error(logServicePath + funcName + " finished with error")
		} else {
			log.Info(logServicePath + funcName + " finished correctly")
		}
	}()
	log.Info(logServicePath + funcName + " started...")

	repUsers, err := s.rep.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	u = user.ConvRepUsersToServUsers(repUsers)
	return u, nil
}

func (s *service) GetUser(ctx context.Context, id string) (u *User, err error) {
	funcName := "GetUser"
	log := logging.GetLogger()
	defer func() {
		if err != nil {
			log.With(err).Error(logServicePath + funcName + " finished with error")
		} else {
			log.Info(logServicePath + funcName + " finished correctly")
		}
	}()
	log.Info(logServicePath + funcName + " started...")

	repUser, err := s.rep.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	u = repUser.ConvToServ()
	return u, nil
}

func (s *service) GetUserByEmail(ctx context.Context, email string) (u *User, err error) {
	funcName := "GetUserByEmail"
	log := logging.GetLogger()
	defer func() {
		if err != nil {
			log.With(err).Error(logServicePath + funcName + " finished with error")
		} else {
			log.Info(logServicePath + funcName + " finished correctly")
		}
	}()
	log.Info(logServicePath + funcName + " started...")

	repUser, err := s.rep.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	u = repUser.ConvToServ()
	return u, nil
}

func (s *service) CreateUser(ctx context.Context, userDto *CreateUserDTO) (id string, err error) {
	funcName := "CreateUser"
	log := logging.GetLogger()
	defer func() {
		if err != nil {
			log.With(err).Error(logServicePath + funcName + " finished with error")
		} else {
			log.Info(logServicePath + funcName + " finished correctly")
		}
	}()
	log.Info(logServicePath + funcName + " started...")

	err = s.validator.Struct(userDto)

	id, err = s.rep.CreateUser(ctx, userDto.ConvToRep())
	if err != nil {
		return id, err
	}

	return id, nil
}

func (s *service) UpdateUser(ctx context.Context, userDto *UpdateUserDTO) (err error) {
	funcName := "UpdateUser"
	log := logging.GetLogger()
	defer func() {
		if err != nil {
			log.With(err).Error(logServicePath + funcName + " finished with error")
		} else {
			log.Info(logServicePath + funcName + " finished correctly")
		}
	}()
	log.Info(logServicePath + funcName + " started...")

	err = s.validator.Struct(userDto)

	err = s.rep.UpdateUser(ctx, userDto.ConvToRep())
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteUser(ctx context.Context, id string) (err error) {
	funcName := "DeleteUser"
	log := logging.GetLogger()
	defer func() {
		if err != nil {
			log.With(err).Error(logServicePath + funcName + " finished with error")
		} else {
			log.Info(logServicePath + funcName + " finished correctly")
		}
	}()
	log.Info(logServicePath + funcName + " started...")

	err = s.rep.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
