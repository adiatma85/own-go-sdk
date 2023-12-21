package jwtAuth

import "github.com/golang-jwt/jwt/v5"

type User struct {
	ID       int64  `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claim struct {
	UserID    int64
	TokenType string
	jwt.RegisteredClaims
}

type UserAuthParam struct {
	User User `json:"user"`
}

type UserAuthInfo struct {
	User User `json:"user"`
}
