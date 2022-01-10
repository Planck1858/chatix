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
	ErrEmptyUser               = errors.New("empty user")
)

const (
	defaultTokenExpiration = time.Minute * 60

	logServicePath = "jwt.helper."
)

type Helper interface {
	GenerateAccessToken(u *user.User) (*Token, error)
	UpdateRefreshToken(rt RefreshToken) (*Token, error)
}

type UserClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

type RefreshToken struct {
	Token string `json:"refresh_token"`
}

type Token struct {
	Token *jwt.Token
	RT    string
}

type helper struct {
	RefreshTokenCache sync.Map // map[string][]byte -> refreshTokenId : []byte(user.User)
}

func NewHelper() Helper {
	return &helper{
		RefreshTokenCache: sync.Map{},
	}
}

func (h *helper) UpdateRefreshToken(rt RefreshToken) (_ *Token, err error) {
	funcName := "UpdateRefreshToken"
	log := logging.GetLogger()
	defer func() {
		if err != nil {
			log.With(err).Error(logServicePath + funcName + " finished with error")
		} else {
			log.Info(logServicePath + funcName + " finished correctly")
		}
	}()
	log.Info(logServicePath + funcName + " started...")

	userBytesRaw, ok := h.RefreshTokenCache.LoadAndDelete(rt.Token)
	if !ok {
		return nil, ErrRefreshTokenDoesntExist
	}

	u := user.User{}
	userBytes, ok := userBytesRaw.([]byte)
	if !ok {
		return nil, ErrInvalidUserBytes
	}
	err = json.Unmarshal(userBytes, &u)
	if err != nil {
		return nil, err
	}

	return h.GenerateAccessToken(&u)
}

func (h *helper) GenerateAccessToken(u *user.User) (t *Token, err error) {
	funcName := "GenerateAccessToken"
	log := logging.GetLogger()
	defer func() {
		if err != nil {
			log.With(err).Error(logServicePath + funcName + " finished with error")
		} else {
			log.Info(logServicePath + funcName + " finished correctly")
		}
	}()
	log.Info(logServicePath + funcName + " started...")

	if u == nil {
		return t, ErrEmptyUser
	}

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
	t.Token, err = builder.Build(claims)
	if err != nil {
		return t, err
	}

	refreshTokenUuid := uuid.New()
	userBytes, _ := json.Marshal(u)
	h.RefreshTokenCache.Store(refreshTokenUuid.String(), userBytes)
	if err != nil {
		return t, err
	}

	t.RT = refreshTokenUuid

	return t, nil
}
