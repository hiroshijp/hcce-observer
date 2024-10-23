package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

var signingKey = []byte(os.Getenv("JWT_SIGNING_KEY"))

func NewJWTMiddleware(e *echo.Group) {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &jwtCustomClaims{}
		},
		SigningKey: signingKey,
	}
	e.Use(echojwt.WithConfig(config))
}

func CreateToken(name string) (string, error) {
	claims := &jwtCustomClaims{
		name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}
