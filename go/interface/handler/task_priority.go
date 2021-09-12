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

type TaskPriorityHandler interface {
	Get() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

// taskPriorityHandler 依存関係
type taskPriorityHandler struct {
	tpu usecase.TaskPriorityUsecase
	ru  usecase.RoleUsecase
}

// NewTaskPriorityHandler 新しくTaskPriorityのハンドラーを作成する。
func NewTaskPriorityHandler(tpu usecase.TaskPriorityUsecase, ru usecase.RoleUsecase) TaskPriorityHandler {
	return &taskPriorityHandler{tpu: tpu, ru: ru}
}

// Get taskPriority全て取得
func (tph *taskPriorityHandler) Get() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if ok := tph.ru.HasRole(c.GetAuthorID(), constant.GetOperation); !ok {
			return c.HasNotRoleResponse()
		}
		taskPriority, err := tph.tpu.Get()
		if err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		data := format.TaskPrioritiesFormat(taskPriority)
		return c.CustomResponse(http.StatusOK, data)
	})
}

func (tph *taskPriorityHandler) GetDetail() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if ok := tph.ru.HasRole(c.GetAuthorID(), constant.GetOperation); !ok {
			return c.HasNotRoleResponse()
		}
		id := tph.getParamID(c)
		taskPriority, err := tph.tpu.GetDetail(id)
		if err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		data := format.TaskPriorityDetailFormat(taskPriority)
		return c.CustomResponse(http.StatusOK, data)
	})
}

func (tph *taskPriorityHandler) Create() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if ok := tph.ru.HasRole(c.GetAuthorID(), constant.CreateOperation); !ok {
			return c.HasNotRoleResponse()
		}
		request := &validation.TaskPriorityRequest{}
		if ok, message := c.BindValidate(request, validation.TaskPriorityMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}
		if err := tph.tpu.Create(request); err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"success": "成功",
		})
	})
}

func (tph *taskPriorityHandler) Update() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if ok := tph.ru.HasRole(c.GetAuthorID(), constant.UpdateOperation); !ok {
			return c.HasNotRoleResponse()
		}
		request := &validation.TaskPriorityRequest{}
		if ok, message := c.BindValidate(request, validation.TaskPriorityMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}
		id := tph.getParamID(c)
		if err := tph.tpu.Update(id, request); err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"success": "成功",
		})
	})
}

func (tph *taskPriorityHandler) Delete() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if ok := tph.ru.HasRole(c.GetAuthorID(), constant.DeleteOperation); !ok {
			return c.HasNotRoleResponse()
		}
		id := tph.getParamID(c)
		if err := tph.tpu.Delete(id); err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"success": "成功",
		})
	})
}

// getParamID URLからidを取得する
func (tph *taskPriorityHandler) getParamID(c *context.CustomContext) uint {
	tmpId, _ := strconv.Atoi(c.Param("id"))
	id := uint(tmpId)
	return id
}
