package handler

import (
	"net/http"
	"wing/application/usecase"
	"wing/interface/context"
	"wing/interface/validation"
	"wing/utils/constant"

	"github.com/labstack/echo/v4"
)

type UserHanlder interface {
	Login() echo.HandlerFunc
	Register() echo.HandlerFunc
	Logout() echo.HandlerFunc
	Edit() echo.HandlerFunc
}

type userHandler struct {
	uu usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) UserHanlder {
	return &userHandler{
		uu: uu,
	}
}

func (uh *userHandler) Login() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		request := &validation.UserLoginRequest{}
		if ok, message := c.BindValidate(request, validation.UserLoginMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}

		token, err := uh.uu.Login(request)
		if err != nil {
			return c.CustomResponse(http.StatusBadRequest, err.Error())
		}

		c.SetCookieToken(token.AccessToken, token.RefreshToken)

		return c.CustomResponse(http.StatusOK, map[string]string{
			"success": "成功",
		})
	})
}

func (uh *userHandler) Register() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		request := &validation.UserRegisterRequest{}
		if ok, message := c.BindValidate(request, validation.UserRegisterMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}

		token, err := uh.uu.Register(request)
		if err != nil {
			return c.CustomResponse(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		c.SetCookieToken(token.AccessToken, token.RefreshToken)

		return c.CustomResponse(http.StatusOK, map[string]string{
			"success": "成功",
		})
	})
}

func (uh *userHandler) Logout() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if err := uh.uu.Logout(c.GetCookieToken()); err != nil {
			return c.CustomResponse(http.StatusBadRequest, err.Error())
		}

		c.DeleteCookie(constant.AccessTokenName)
		c.DeleteCookie(constant.RefreshTokenName)

		return c.CustomResponse(http.StatusOK, map[string]string{
			"success": "ログアウト成功",
		})
	})
}

func (uh *userHandler) Edit() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		request := &validation.UserEditRequest{}
		if ok, message := c.BindValidate(request, validation.UserEditMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}

		user, err := uh.uu.Edit(request, c.GetAuthorID())
		if err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}

		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"user": map[string]interface{}{
				"ID":              user.ID,
				"name":            user.Name,
				"Email":           user.Email,
				"EmailVerifiedAt": user.EmailVerifiedAt,
			},
		})
	})
}
