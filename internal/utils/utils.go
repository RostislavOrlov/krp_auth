package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"krp_project/internal/dto"
	"krp_project/internal/entities"
	"strconv"
	"time"
)

func GeneratePassword(user *dto.RegisterRequest) string {
	return user.FirstName
}

func CreateTokens(usr *entities.User) ([]*entities.Token, error) {
	mySigningKey := []byte("SecretKey")
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

type TokenClaims struct {
	jwt.StandardClaims
}

func UpdateTokens(tokens *entities.Tokens) ([]*entities.Token, error) {

	mySigningKey := []byte("SecretKey")

	accessTokenWithClaims, err := jwt.ParseWithClaims(tokens.AccessTokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claimsAccess, okAccess := accessTokenWithClaims.Claims.(*TokenClaims); okAccess && accessTokenWithClaims.Valid {
		fmt.Printf("%v %v", claimsAccess.StandardClaims.ExpiresAt)
		accessTokenClaims := jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			//Subject:   usr.LastName + usr.FirstName + usr.MiddleName + strconv.Itoa(usr.Id),
		}

		refreshTokenClaims := jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour * 31).Unix(),
			IssuedAt:  time.Now().Unix(),
			//Subject:   usr.LastName + usr.FirstName + usr.MiddleName + strconv.Itoa(usr.Id),
		}

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
	} else {
		fmt.Println(err.Error())
	}

	return nil, err
}
