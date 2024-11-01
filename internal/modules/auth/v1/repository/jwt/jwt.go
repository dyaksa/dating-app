package jwt

import (
	"context"
	"os"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/dyaksa/dating-app/internal/modules/auth/v1/domain"
)

type RepositoryJWT struct{}

func NewJWTRepository() *RepositoryJWT {
	return &RepositoryJWT{}
}

func (r *RepositoryJWT) Generate(ctx context.Context, payload *domain.TokenClaim, exipred time.Duration) (string, error) {
	var token = new(jwtgo.Token)
	exp := time.Now().Add(exipred)
	now := time.Now()

	var key interface{}
	token = jwtgo.New(jwtgo.SigningMethodHS256)
	key = []byte(os.Getenv("SECRET_KEY"))

	claims := jwtgo.MapClaims{
		"iss": "dating-app",
		"exp": exp.Unix(),
		"iat": now.Unix(),
		"sub": payload.User.ID,
	}

	token.Claims = claims

	tokenString, err := token.SignedString(key.([]byte))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
