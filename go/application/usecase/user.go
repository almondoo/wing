package usecase

import (
	"wing/domain/entity"
	"wing/domain/service"
	"wing/infrastructure/auth"
	"wing/interface/validation"
)

type UserUsecase interface {
	Login(*validation.UserLoginRequest) (*auth.TokenDetails, error)
	Register(*validation.UserRegisterRequest) (*auth.TokenDetails, error)
	Logout(*auth.TokenData) error
	Edit(*validation.UserEditRequest, uint) (*entity.User, error)
	EditAssignRole(*validation.UserEditAssignRoleRequest) error
}

type userUsecase struct {
	us    service.UserService
	rs    service.RoleService
	auth  auth.AuthInterface
	token auth.TokenInterface
}

func NewUserUsecase(us service.UserService, rs service.RoleService, auth auth.AuthInterface, token auth.TokenInterface) UserUsecase {
	return &userUsecase{us: us, rs: rs, auth: auth, token: token}
}

func (uu *userUsecase) Login(request *validation.UserLoginRequest) (*auth.TokenDetails, error) {
	user, err := uu.us.Login(request)
	if err != nil {
		return nil, err
	}

	return uu.createToken(int(user.ID))
}

func (uu *userUsecase) Register(request *validation.UserRegisterRequest) (*auth.TokenDetails, error) {
	user, err := uu.us.Register(request)
	if err != nil {
		return nil, err
	}

	return uu.createToken(int(user.ID))
}

func (uu *userUsecase) Logout(tokens *auth.TokenData) error {
	if err := uu.us.Logout(tokens); err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) Edit(request *validation.UserEditRequest, userID uint) (*entity.User, error) {
	user, err := uu.us.Edit(request, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu *userUsecase) EditAssignRole(*validation.UserEditAssignRoleRequest) error {
	// uu.us.Edit()
	return nil
}

func (uu *userUsecase) createToken(id int) (*auth.TokenDetails, error) {
	token, err := uu.token.CreateToken(id)
	if err != nil {
		return nil, err
	}

	if err := uu.auth.CreateAuth(id, token); err != nil {
		return nil, err
	}

	return token, nil
}
