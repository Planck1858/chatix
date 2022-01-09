package jwt

import (
	"encoding/json"
	"githab.com/Planck1858/chatix/back-end/internal/config"
	"githab.com/Planck1858/chatix/back-end/internal/user"
	"githab.com/Planck1858/chatix/back-end/pkg/logging"
	"github.com/cristalhq/jwt/v3"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"sync"
	"time"
)

var (
	ErrRefreshTokenDoesntExist = errors.New("refresh token doesn't exist")
	ErrInvalidUserBytes        = errors.New("invalid user bytes")
)

const (
	defaultTokenExpiration = time.Minute * 60
)

type Helper interface {
	GenerateAccessToken(u user.User) (string, error)
	UpdateRefreshToken(rt RefreshToken) (string, error)
}

type helper struct {
	Log               logging.Logger
	RefreshTokenCache sync.Map // map[string][]byte -> refreshTokenUuid : user.User([]byte)
}

func NewHelper(logger logging.Logger) Helper {
	return &helper{
		Log:               logger,
		RefreshTokenCache: sync.Map{},
	}
}

type UserClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

func (h *helper) UpdateRefreshToken(rt RefreshToken) (string, error) {
	userBytesRaw, ok := h.RefreshTokenCache.LoadAndDelete(rt.RefreshToken)
	if !ok {
		return "", ErrRefreshTokenDoesntExist
	}

	u := user.User{}
	userBytes, ok := userBytesRaw.([]byte)
	if !ok {
		return "", ErrInvalidUserBytes
	}
	err := json.Unmarshal(userBytes, &u)
	if err != nil {
		return "", err
	}

	return h.GenerateAccessToken(u)
}

func (h *helper) GenerateAccessToken(u user.User) (t string, err error) {
	h.Log.Info("jwt.GenerateAccessToken starting")

	defer func() {
		if err != nil {
			h.Log.Error(errors.Wrap(err, "jwt.GenerateAccessToken error"))
		}
	}()

	secretKey := []byte(config.GetConfig().Auth.Secret)
	signer, err := jwt.NewSignerHS(jwt.HS256, secretKey)
	if err != nil {
		return t, err
	}
	builder := jwt.NewBuilder(signer)

	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        u.Id,
			Audience:  []string{u.Role.String()},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(defaultTokenExpiration)),
		},
		Email: u.Email,
	}
	token, err := builder.Build(claims)
	if err != nil {
		return t, err
	}

	refreshTokenUuid := uuid.New()
	userBytes, _ := json.Marshal(u)
	h.RefreshTokenCache.Store(refreshTokenUuid.String(), userBytes)
	if err != nil {
		return t, err
	}

	jsonBytes, err := json.Marshal(map[string]string{
		"token":         token.String(),
		"refresh_token": refreshTokenUuid.String(),
	})
	if err != nil {
		return t, err
	}
	t = string(jsonBytes)

	h.Log.Info("jwt.GenerateAccessToken finished correctly")
	return t, nil
}
