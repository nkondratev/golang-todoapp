package domain

import "time"

type JWTToken struct {
	Name  string
	Token string
	Exp   time.Time
}

func NewJWTToken(token string, exp time.Time) JWTToken {
	return JWTToken{
		Token: token,
		Exp:   exp,
		Name:  "jwt",
	}
}
