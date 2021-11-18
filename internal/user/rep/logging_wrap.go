package rep

import (
	"context"
	"githab.com/Planck1858/chatix/pkg/logging"
	"github.com/davecgh/go-spew/spew"
)

type loggingRep struct {
	next Repository
}

func RepositoryWithLogger(next Repository) Repository {
	return &loggingRep{
		next: next,
	}
}

func (s *loggingRep) GetAllUsers(ctx context.Context) (_ []User, err error) {
	log, _ := logging.GetFromContext(ctx)
	log.Info("user.service.GetAllUsers started")

	defer func() {
		if err != nil {
			log.Error("user.service.GetAllUsers failed: ", err)
		} else {
			log.Info("user.service.GetAllUsers succeed")
		}
	}()

	return s.next.GetAllUsers(ctx)
}

func (s *loggingRep) GetUser(ctx context.Context, id string) (_ *User, err error) {
	log, _ := logging.GetFromContext(ctx)
	log.With("id", id).
		Info("user.service.GetUser started")

	defer func() {
		if err != nil {
			log.Error("user.service.GetUser failed: ", err)
		} else {
			log.Info("user.service.GetUser succeed")
		}
	}()

	return s.next.GetUser(ctx, id)
}

func (s *loggingRep) CreateUser(ctx context.Context, user *User) (id string, err error) {
	log, _ := logging.GetFromContext(ctx)
	log.With("user", spew.Sprintln(user)).
		Info("user.service.CreateUser started")

	defer func() {
		if err != nil {
			log.Error("user.service.CreateUser failed: ", err)
		} else {
			log.With("user id", id).
				Info("user.service.CreateUser succeed")
		}
	}()

	return s.next.CreateUser(ctx, user)
}

func (s *loggingRep) UpdateUser(ctx context.Context, user *User) (err error) {
	log, _ := logging.GetFromContext(ctx)
	log.With("user", spew.Sprintln(user)).
		Info("user.service.UpdateUser started")

	defer func() {
		if err != nil {
			log.Error("user.service.UpdateUser failed: ", err)
		} else {
			log.Info("user.service.UpdateUser succeed")
		}
	}()

	return s.next.UpdateUser(ctx, user)
}

func (s *loggingRep) DeleteUser(ctx context.Context, id string) (err error) {
	log, _ := logging.GetFromContext(ctx)
	log.With("id", id).
		Info("user.service.DeleteUser started")

	defer func() {
		if err != nil {
			log.Error("user.service.DeleteUser failed: ", err)
		} else {
			log.Info("user.service.DeleteUser succeed")
		}
	}()

	return s.next.DeleteUser(ctx, id)
}
