package services

import (
	"errors"
	"krp_project/internal/dto"
	"krp_project/internal/entities"
	"krp_project/internal/repositories"
	"krp_project/internal/utils"
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
	password := utils.GeneratePassword(user)
	usr, err := srv.repo.Register(user, password)
	if err != nil {
		return nil, errors.New("failed user registration")
	}

	return usr, nil
}

func (srv *UserService) Auth(user *dto.AuthRequest) (*entities.User, []*entities.Token, error) {
	usr, err := srv.repo.Auth(user)
	if err != nil {
		return nil, nil, errors.New("error user authentication")
	}
	if user.Password != usr.Password {
		return nil, nil, errors.New("incorrect password")
	}

	userTokens, _ := utils.UpdateTokens(usr)

	//srv.repo.UpdateRefreshToken?
	return usr, userTokens, nil
}

func (srv *UserService) UpdateAccessToken(req *dto.UpdateAccessTokenRequest) ([]*entities.Token, error) {
	usr := &entities.User{
		Id:         req.Id,
		LastName:   req.LastName,
		FirstName:  req.FirstName,
		MiddleName: req.MiddleName,
		Email:      req.Email,
		Role:       req.Role,
	}

	return utils.UpdateTokens(usr)
}
