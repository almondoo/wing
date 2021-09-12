package handler

import (
	"net/http"
	"strconv"
	"wing/application/format"
	"wing/application/usecase"
	"wing/interface/context"
	"wing/interface/validation"
	"wing/utils/constant"

	"github.com/labstack/echo/v4"
)

type TaskChildHandler interface {
	Get() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

// taskChildHandler 依存関係
type taskChildHandler struct {
	tcu usecase.TaskChildUsecase
	ru  usecase.RoleUsecase
}

// NewTaskChildHandler 新しくTaskChildのハンドラーを作成する。
func NewTaskChildHandler(tcu usecase.TaskChildUsecase, ru usecase.RoleUsecase) TaskChildHandler {
	return &taskChildHandler{tcu: tcu, ru: ru}
}

// Get task status全て取得
func (tch *taskChildHandler) Get() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if ok := tch.ru.HasRole(c.GetAuthorID(), constant.GetOperation); !ok {
			return c.HasNotRoleResponse()
		}
		taskChildren, err := tch.tcu.Get()
		if err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		data := format.TaskChildrenFormat(taskChildren)
		return c.CustomResponse(http.StatusOK, data)
	})
}

// GetDetail 詳細取得
func (tch *taskChildHandler) GetDetail() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if ok := tch.ru.HasRole(c.GetAuthorID(), constant.GetOperation); !ok {
			return c.HasNotRoleResponse()
		}
		id := tch.getParamID(c)
		taskChild, err := tch.tcu.GetDetail(id)
		if err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		data := format.TaskChildDetailFormat(taskChild)
		return c.CustomResponse(http.StatusOK, data)
	})
}

// Create 作成
func (tch *taskChildHandler) Create() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if ok := tch.ru.HasRole(c.GetAuthorID(), constant.CreateOperation); !ok {
			return c.HasNotRoleResponse()
		}
		request := &validation.TaskChildRequest{}
		if ok, message := c.BindValidate(request, validation.TaskChildMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}
		if err := tch.tcu.Create(request); err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"success": "成功",
		})
	})
}

// Update 更新
func (tch *taskChildHandler) Update() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if ok := tch.ru.HasRole(c.GetAuthorID(), constant.UpdateOperation); !ok {
			return c.HasNotRoleResponse()
		}
		request := &validation.TaskChildRequest{}
		if ok, message := c.BindValidate(request, validation.TaskChildMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}
		id := tch.getParamID(c)
		if err := tch.tcu.Update(id, request); err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"success": "成功",
		})
	})
}

// Delete 削除
func (tch *taskChildHandler) Delete() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if ok := tch.ru.HasRole(c.GetAuthorID(), constant.DeleteOperation); !ok {
			return c.HasNotRoleResponse()
		}
		id := tch.getParamID(c)
		if err := tch.tcu.Delete(id); err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"success": "成功",
		})
	})
}

// getParamID URLからidを取得する
func (tch *taskChildHandler) getParamID(c *context.CustomContext) uint {
	tmpId, _ := strconv.Atoi(c.Param("id"))
	id := uint(tmpId)
	return id
}
