package utils

import (
	"github.com/golang-jwt/jwt"
	"krp_project/internal/dto"
	"krp_project/internal/entities"
	"strconv"
	"time"
)

func GeneratePassword(user *dto.RegisterRequest) string {
	return user.FirstName
}

func UpdateTokens(usr *entities.User) ([]*entities.Token, error) {
	accessTokenClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   usr.LastName + usr.FirstName + usr.MiddleName + strconv.Itoa(usr.Id),
	}

	refreshTokenClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour * 31).Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   usr.LastName + usr.FirstName + usr.MiddleName + strconv.Itoa(usr.Id),
	}

	mySigningKey := []byte("SecretKey")

	accessTokenTemp := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	signedAccessToken, _ := accessTokenTemp.SignedString(mySigningKey)
	accessToken := entities.Token{
		TokenString: signedAccessToken,
		ExpiresAt:   accessTokenClaims.ExpiresAt,
		IssuedAt:    accessTokenClaims.IssuedAt,
	}

	refreshTokenTemp := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	signedRefreshToken, _ := refreshTokenTemp.SignedString(mySigningKey)
	refreshToken := entities.Token{
		TokenString: signedRefreshToken,
		ExpiresAt:   accessTokenClaims.ExpiresAt,
		IssuedAt:    accessTokenClaims.IssuedAt,
	}

	return []*entities.Token{&accessToken, &refreshToken}, nil
}
