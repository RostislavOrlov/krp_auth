package services

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"krp_project/internal/dto"
	"krp_project/internal/entities"
	"krp_project/internal/repositories"
	"strconv"
	"time"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) (*UserService, error) {
	return &UserService{
		repo: repo,
	}, nil
}

func (srv *UserService) Register(user *dto.RegisterRequest) (*entities.User, error) {
	usr, err := srv.repo.Register(user)
	if err != nil {
		return nil, errors.New("failed user registration")
	}

	return usr, nil
}

func (srv *UserService) Auth(user *dto.AuthRequest) (*entities.User, []*entities.Token, error) {
	usr, err := srv.repo.Auth(user)
	if err != nil {

	}
	if user.Password != usr.Password {
		return nil, nil, errors.New("incorrect password")
	}

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
	signedAccessToken, err := accessTokenTemp.SignedString(mySigningKey)
	accessToken := entities.Token{
		TokenString: signedAccessToken,
		ExpiresAt:   accessTokenClaims.ExpiresAt,
		IssuedAt:    accessTokenClaims.IssuedAt,
	}

	refreshTokenTemp := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	signedRefreshToken, err := refreshTokenTemp.SignedString(mySigningKey)
	refreshToken := entities.Token{
		TokenString: signedRefreshToken,
		ExpiresAt:   accessTokenClaims.ExpiresAt,
		IssuedAt:    accessTokenClaims.IssuedAt,
	}

	//srv.repo.UpdateRefreshToken?
	return usr, []*entities.Token{&accessToken, &refreshToken}, nil
}
