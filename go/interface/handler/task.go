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

type TaskHandler interface {
	Get() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

// taskHandler 依存関係
type taskHandler struct {
	tu usecase.TaskUsecase
	ru usecase.RoleUsecase
}

// NewTaskHandler 新しくTaskのハンドラーを作成する。
func NewTaskHandler(tu usecase.TaskUsecase, ru usecase.RoleUsecase) TaskHandler {
	return &taskHandler{tu: tu, ru: ru}
}

// Get task全て取得
func (th *taskHandler) Get() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if ok := th.ru.HasRole(c.GetAuthorID(), constant.GetOperation); !ok {
			return c.HasNotRoleResponse()
		}
		tasks, err := th.tu.Get()
		if err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		data := format.TasksFormat(tasks)
		return c.CustomResponse(http.StatusOK, data)
	})
}

func (th *taskHandler) GetDetail() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if ok := th.ru.HasRole(c.GetAuthorID(), constant.GetOperation); !ok {
			return c.HasNotRoleResponse()
		}
		id := th.getParamID(c)
		task, err := th.tu.GetDetail(id)
		if err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		data := format.TaskDetailFormat(task)
		return c.CustomResponse(http.StatusOK, data)
	})
}

func (th *taskHandler) Create() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if ok := th.ru.HasRole(c.GetAuthorID(), constant.CreateOperation); !ok {
			return c.HasNotRoleResponse()
		}
		request := &validation.TaskRequest{}
		if ok, message := c.BindValidate(request, validation.TaskMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}
		if err := th.tu.Create(request); err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"success": "成功",
		})
	})
}

func (th *taskHandler) Update() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if ok := th.ru.HasRole(c.GetAuthorID(), constant.UpdateOperation); !ok {
			return c.HasNotRoleResponse()
		}
		request := &validation.TaskRequest{}
		if ok, message := c.BindValidate(request, validation.TaskMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}
		id := th.getParamID(c)
		if err := th.tu.Update(id, request); err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"success": "成功",
		})
	})
}

func (th *taskHandler) Delete() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if ok := th.ru.HasRole(c.GetAuthorID(), constant.DeleteOperation); !ok {
			return c.HasNotRoleResponse()
		}
		id := th.getParamID(c)
		if err := th.tu.Delete(id); err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"success": "成功",
		})
	})
}

// getParamID URLからidを取得する
func (th *taskHandler) getParamID(c *context.CustomContext) uint {
	tmpId, _ := strconv.Atoi(c.Param("id"))
	id := uint(tmpId)
	return id
}
