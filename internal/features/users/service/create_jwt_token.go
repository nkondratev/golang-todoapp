package users_service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nkondratev/golang-todoapp/internal/core/domain"
)

func (u *UsersService) CreateJWTToken(user domain.User) (domain.JWTToken, error) {
	exp := time.Now().Add(time.Hour * 24)
	key := os.Getenv("JWT_KEY")
	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": exp.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims, nil)
	jwtToken, err := token.SignedString([]byte(key))
	return domain.NewJWTToken(jwtToken, exp), err
}
