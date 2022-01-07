package jwt

import (
	"context"
	"encoding/json"
	"githab.com/Planck1858/chatix/internal/config"
	"github.com/cristalhq/jwt/v3"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

const UserUUID = "user_uuid"

func Middleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("jwt.Middleware starting")

		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			unauthorized(w, errors.New("incorrect auth header"))
			return
		}

		secretKey := []byte(config.GetConfig().Auth.Secret)
		jwtToken := authHeader[1]
		verifier, err := jwt.NewVerifierHS(jwt.HS256, secretKey)
		if err != nil {
			unauthorized(w, err)
			return
		}
		token, err := jwt.ParseAndVerifyString(jwtToken, verifier)
		if err != nil {
			unauthorized(w, err)
			return
		}

		var uc UserClaims
		err = json.Unmarshal(token.RawClaims(), &uc)
		if err != nil {
			unauthorized(w, err)
			return
		}
		if valid := uc.IsValidAt(time.Now()); !valid {
			unauthorized(w, err)
			return
		}

		ctx := context.WithValue(r.Context(), UserUUID, uc.ID)
		h(w, r.WithContext(ctx))
		logrus.Info("jwt.Middleware finished correctly")
	}
}

func unauthorized(w http.ResponseWriter, err error) {
	logrus.Errorf("jwt.Middleware error: %v", err)
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("unauthorized"))
}
