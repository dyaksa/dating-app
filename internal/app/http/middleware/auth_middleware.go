package middleware

import (
	"errors"
	"os"
	"strings"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/dyaksa/dating-app/internal/modules/auth/v1/domain"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct{}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		authValues := strings.Split(authorization, " ")
		authType := strings.ToLower(authValues[0])
		if authType != "bearer" || len(authValues) != 2 {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		tokenString := authValues[1]
		tokenParse, err := jwtgo.Parse(tokenString, func(token *jwtgo.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		var errToken error
		switch ve := err.(type) {
		case *jwtgo.ValidationError:
			if ve.Errors == jwtgo.ValidationErrorExpired {
				errToken = errors.New("token expired")
			} else {
				errToken = errors.New("token format")
			}
		}

		if errToken != nil {
			c.JSON(401, gin.H{"message": errToken.Error()})
			c.Abort()
			return
		}

		if !tokenParse.Valid {
			c.JSON(401, gin.H{"message": "token is not valid"})
			c.Abort()
		}

		mapClaims, _ := tokenParse.Claims.(jwtgo.MapClaims)

		var tokenClaim domain.TokenClaim
		tokenClaim.User.ID = mapClaims["sub"].(string)

		c.Set("user_id", tokenClaim.User.ID)
		c.Next()
	}
}
