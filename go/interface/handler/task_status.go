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

type TaskStatusHandler interface {
	Get() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

// taskStatusHandler 依存関係
type taskStatusHandler struct {
	tsu usecase.TaskStatusUsecase
}

// NewTaskStatusHandler 新しくTaskStatusのハンドラーを作成する。
func NewTaskStatusHandler(tsu usecase.TaskStatusUsecase) TaskStatusHandler {
	return &taskStatusHandler{tsu: tsu}
}

// Get task status全て取得
func (tsh *taskStatusHandler) Get() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		taskStatuses, err := tsh.tsu.Get()
		if err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		data := format.TaskStatusesFormat(taskStatuses)
		return c.CustomResponse(http.StatusOK, data)
	})
}

// GetDetail 詳細取得
func (tsh *taskStatusHandler) GetDetail() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		id := tsh.getParamID(c)
		taskStatus, err := tsh.tsu.GetDetail(id)
		if err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		data := format.TaskStatusDetailFormat(taskStatus)
		return c.CustomResponse(http.StatusOK, data)
	})
}

// Create 作成
func (tsh *taskStatusHandler) Create() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		request := &validation.TaskStatusRequest{}
		if ok, message := c.BindValidate(request, validation.TaskStatusMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}
		if err := tsh.tsu.Create(request); err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"success": "成功",
		})
	})
}

// Update 更新
func (tsh *taskStatusHandler) Update() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		request := &validation.TaskStatusRequest{}
		if ok, message := c.BindValidate(request, validation.TaskStatusMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}
		id := tsh.getParamID(c)
		if err := tsh.tsu.Update(id, request); err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"success": "成功",
		})
	})
}

// Delete 削除
func (tsh *taskStatusHandler) Delete() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		id := tsh.getParamID(c)
		if err := tsh.tsu.Delete(id); err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"success": "成功",
		})
	})
}

// getParamID URLからidを取得する
func (tsh *taskStatusHandler) getParamID(c *context.CustomContext) uint {
	tmpId, _ := strconv.Atoi(c.Param("id"))
	id := uint(tmpId)
	return id
}
