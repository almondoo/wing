package handler

import (
	"net/http"
	"strconv"
	"wing/application/format"
	"wing/application/usecase"
	"wing/interface/context"
	"wing/interface/validation"

	"github.com/labstack/echo/v4"
)

type RoleHandler interface {
	Get() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

// roleHandler 依存関係
type roleHandler struct {
	ru usecase.RoleUsecase
}

// NewRoleHandler 新しくRoleのハンドラーを作成する。
func NewRoleHandler(ru usecase.RoleUsecase) RoleHandler {
	return &roleHandler{ru: ru}
}

// Get Role全て取得
func (rh *roleHandler) Get() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		roles, err := rh.ru.Get()
		if err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		data := format.RolesFormat(roles)
		return c.CustomResponse(http.StatusOK, data)
	})
}

// GetDetail Role詳細取得
func (rh *roleHandler) GetDetail() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		id := rh.getParamID(c)
		role, err := rh.ru.GetDetail(id)
		if err != nil {
			return c.CustomResponse(http.StatusBadRequest, err.Error())
		}
		data := format.RoleDetailFormat(role)

		return c.CustomResponse(http.StatusOK, data)
	})
}

// Create Roleを作成する
func (rh *roleHandler) Create() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		request := &validation.RoleRequest{}
		if ok, message := c.BindValidate(request, validation.RoleMessage); !ok {
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

// Update Roleを更新する
func (rh *roleHandler) Update() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		request := &validation.RoleRequest{}
		if ok, message := c.BindValidate(request, validation.RoleMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}
		id := rh.getParamID(c)
		if err := rh.ru.Update(id, request); err != nil {
			return c.CustomResponse(http.StatusBadRequest, err.Error())
		}
		return c.CustomResponse(http.StatusOK, map[string]string{
			"success": "ok",
		})
	})
}

// Delete Roleを削除する
func (rh *roleHandler) Delete() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		id := rh.getParamID(c)
		if err := rh.ru.Delete(id); err != nil {
			return c.CustomResponse(http.StatusBadRequest, err.Error())
		}
		return c.CustomResponse(http.StatusOK, map[string]string{
			"success": "ok",
		})
	})
}

// getParamID URLからidを取得する
func (rh *roleHandler) getParamID(c *context.CustomContext) uint {
	tmpId, _ := strconv.Atoi(c.Param("id"))
	id := uint(tmpId)
	return id
}
