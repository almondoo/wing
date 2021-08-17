package service

import (
	"wing/domain/entity"
	"wing/domain/repository"
	"wing/infrastructure/auth"
	"wing/infrastructure/security"
	"wing/interface/validation"
)

type UserService interface {
	Login(*validation.UserLoginRequest) (*entity.User, error)
	Register(*validation.UserRegisterRequest) (*entity.User, error)
	Logout(*auth.TokenData) error
	Edit(*validation.UserEditRequest, uint) (*entity.User, error)
}

type userService struct {
	userRepo repository.UserRepository
	auth     auth.AuthInterface
	token    auth.TokenInterface
}

func NewUserService(userRepo repository.UserRepository, auth auth.AuthInterface, token auth.TokenInterface) UserService {
	return &userService{userRepo: userRepo, auth: auth, token: token}
}

func (us *userService) Login(request *validation.UserLoginRequest) (*entity.User, error) {
	user, err := us.userRepo.FindByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	if err := security.VerifyPassword(user.Password, request.Password); err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userService) Register(request *validation.UserRegisterRequest) (*entity.User, error) {
	password, err := security.Hash(request.Password)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: password,
	}

	createUser, err := us.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createUser, nil
}

func (us *userService) Logout(tokens *auth.TokenData) error {
	detailtoken, err := us.token.ExtractTokenMetadata(tokens.AccessToken)
	if err != nil {
		return err
	}
	if err := us.auth.DeleteTokens(detailtoken); err != nil {
		return err
	}
	return nil
}

func (us *userService) Edit(request *validation.UserEditRequest, userID uint) (user *entity.User, err error) {
	user, err = us.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	user.Name = request.Name
	user.Email = request.Email
	user, err = us.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
