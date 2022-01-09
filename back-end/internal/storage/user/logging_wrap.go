package user

import (
	"context"
	"githab.com/Planck1858/chatix/back-end/pkg/logging"
	"github.com/davecgh/go-spew/spew"
)

const logRepPath = "user.service."

type loggingRep struct {
	next Repository
}

func RepositoryWithLogger(next Repository) Repository {
	return &loggingRep{
		next: next,
	}
}

func (s *loggingRep) GetAllUsers(ctx context.Context) (u []User, err error) {
	funcName := "GetAllUsers"
	log := logging.GetLogger()
	log.Info(logRepPath + funcName + " started")

	defer func() {
		if err != nil {
			log.Error(logRepPath+funcName+" failed: ", err)
		} else {
			log.With("users", spew.Sprintln(u)).
				Info(logRepPath + funcName + " succeed")
		}
	}()

	return s.next.GetAllUsers(ctx)
}

func (s *loggingRep) GetUser(ctx context.Context, id string) (u *User, err error) {
	funcName := "GetUser"
	log := logging.GetLogger()
	log.With("id", id).
		Info(logRepPath + funcName + " started")

	defer func() {
		if err != nil {
			log.Error(logRepPath+funcName+" failed: ", err)
		} else {
			log.With("user", spew.Sprintln(u)).
				Info(logRepPath + funcName + " succeed")
		}
	}()

	return s.next.GetUser(ctx, id)
}

func (s *loggingRep) GetUserByEmail(ctx context.Context, email string) (u *User, err error) {
	funcName := "GetUserByEmail"
	log := logging.GetLogger()
	log.With("email", email).
		Info(logRepPath + funcName + " started")

	defer func() {
		if err != nil {
			log.Error(logRepPath+funcName+" failed: ", err)
		} else {
			log.With("user", spew.Sprintln(u)).
				Info(logRepPath + funcName + " succeed")
		}
	}()

	return s.next.GetUserByEmail(ctx, email)
}

func (s *loggingRep) CreateUser(ctx context.Context, user *User) (id string, err error) {
	funcName := "CreateUser"
	log := logging.GetLogger()
	log.With("user", spew.Sprintln(user)).
		Info(logRepPath + funcName + " started")

	defer func() {
		if err != nil {
			log.Error(logRepPath+funcName+" failed: ", err)
		} else {
			log.With("user id", id).
				Info(logRepPath + funcName + " succeed")
		}
	}()

	return s.next.CreateUser(ctx, user)
}

func (s *loggingRep) UpdateUser(ctx context.Context, user *User) (err error) {
	funcName := "UpdateUser"
	log := logging.GetLogger()
	log.With("user", spew.Sprintln(user)).
		Info(logRepPath + "UpdateUser started")

	defer func() {
		if err != nil {
			log.Error(logRepPath+funcName+" failed: ", err)
		} else {
			log.Info(logRepPath + funcName + " succeed")
		}
	}()

	return s.next.UpdateUser(ctx, user)
}

func (s *loggingRep) DeleteUser(ctx context.Context, id string) (err error) {
	funcName := "DeleteUser"
	log := logging.GetLogger()
	log.With("id", id).
		Info(logRepPath + funcName + " started")

	defer func() {
		if err != nil {
			log.Error(logRepPath+funcName+" failed: ", err)
		} else {
			log.Info(logRepPath + funcName + " succeed")
		}
	}()

	return s.next.DeleteUser(ctx, id)
}
