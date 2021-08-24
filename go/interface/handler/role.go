package handler

import (
	"net/http"
	"wing/application/usecase"
	"wing/interface/context"
	"wing/interface/validation"

	"github.com/labstack/echo/v4"
)

type RoleHandler interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
}

// roleHandler 依存関係
type roleHandler struct {
	ru usecase.RoleUsecase
}

// NewRoleHandler 新しく権限のハンドラーを作成する。
func NewRoleHandler(ru usecase.RoleUsecase) RoleHandler {
	return &roleHandler{ru: ru}
}

// Create 権限を作成する
func (rh *roleHandler) Create() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		request := &validation.RoleCreateRequest{}
		if ok, message := c.BindValidate(request, validation.RoleCreateMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}
		if err := rh.ru.Create(request); err != nil {
			return c.CustomResponse(http.StatusBadRequest, err.Error())
		}
		return c.CustomResponse(http.StatusOK, map[string]string{
			"success": "ok",
		})
	})
}

// Update 権限を更新する
func (rh *roleHandler) Update() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		request := &validation.RoleUpdateRequest{}
		if ok, message := c.BindValidate(request, validation.RoleUpdateMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}
		if err := rh.ru.Update(request); err != nil {
			return c.CustomResponse(http.StatusBadRequest, err.Error())
		}
		return c.CustomResponse(http.StatusOK, map[string]string{
			"success": "ok",
		})
	})
}
