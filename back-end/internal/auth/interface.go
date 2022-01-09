package auth

import "context"

type Service interface {
	SignUp(ctx context.Context, dto *SignUpDTO) (userId string, err error)
	SignIn(ctx context.Context, dto *SignInDTO) (ok bool, err error)
}
