package jwtAuth

import (
	"context"
	"time"

	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/own-go-sdk/errors"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	userAuthInfo     contextKey = "UserAuthInfo"
	AccessTokenType             = "ACCESS_TOKEN"
	RefreshTokenType            = "REFRESH_TOKEN"
)

type Interface interface {
	CreateAccessToken(user User) (string, error)
	CreateRefreshToken(user User) (string, error)
	ValidateToken(token string) (User, error)
	SetUserAuthInfo(ctx context.Context, param UserAuthParam) context.Context
	GetUserAuthInfo(ctx context.Context) (UserAuthInfo, error)
}

type Config struct {
	AccessTokenExpLimit  time.Duration
	RefreshTokenExpLimit time.Duration
	Secret               string
	secretByte           []byte
}

type jsonWebtoken struct {
	cfg Config
}

func Init(conf Config) Interface {
	conf.secretByte = []byte(conf.Secret)
	j := &jsonWebtoken{
		cfg: conf,
	}

	return j
}

func (j *jsonWebtoken) CreateAccessToken(user User) (string, error) {
	expirationTime := time.Now().Add(j.cfg.AccessTokenExpLimit)

	claims := &Claim{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.cfg.secretByte)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *jsonWebtoken) CreateRefreshToken(user User) (string, error) {
	expirationTime := time.Now().Add(j.cfg.RefreshTokenExpLimit)

	claims := &Claim{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.cfg.secretByte)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *jsonWebtoken) ValidateToken(token string) (User, error) {
	var (
		codeError    codes.Code
		errorMessage string
		user         User
		claim        Claim
	)

	parsedToken, err := jwt.ParseWithClaims(token, &claim, func(t *jwt.Token) (interface{}, error) {
		return j.cfg.secretByte, nil
	})
	if err != nil {
		return user, err
	}

	switch claim.TokenType {
	case AccessTokenType:
		codeError = codes.CodeAuthAccessTokenExpired
		errorMessage = "access token is not valid"
	case RefreshTokenType:
		codeError = codes.CodeAuthRefreshTokenExpired
		errorMessage = "refresh token is not valid"
	}

	if !parsedToken.Valid {
		return user, errors.NewWithCode(codeError, errorMessage)
	}

	// If anything else is valid
	user.ID = claim.UserID

	return user, nil
}

func (j *jsonWebtoken) SetUserAuthInfo(ctx context.Context, param UserAuthParam) context.Context {
	userauth := UserAuthInfo{
		User: param.User,
	}
	return context.WithValue(ctx, userAuthInfo, userauth)
}

func (j *jsonWebtoken) GetUserAuthInfo(ctx context.Context) (UserAuthInfo, error) {
	user, ok := ctx.Value(userAuthInfo).(UserAuthInfo)
	if !ok {
		return user, errors.NewWithCode(codes.CodeAuthFailure, "failed getting user auth info")
	}

	return user, nil
}
