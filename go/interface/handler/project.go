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

type ProjectHandler interface {
	Get() echo.HandlerFunc
	GetDetail() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

// projectHandler 依存関係
type projectHandler struct {
	pu usecase.ProjectUsecase
	ru usecase.RoleUsecase
}

// NewProjectHandler 新しくProjectのハンドラーを作成する。
func NewProjectHandler(pu usecase.ProjectUsecase, ru usecase.RoleUsecase) ProjectHandler {
	return &projectHandler{pu: pu, ru: ru}
}

// Get task status全て取得
func (ph *projectHandler) Get() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if ok := ph.ru.HasRole(c.GetAuthorID(), constant.GetOperation); !ok {
			return c.HasNotRoleResponse()
		}
		projects, err := ph.pu.Get()
		if err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		data := format.ProjectsFormat(projects)
		return c.CustomResponse(http.StatusOK, data)
	})
}

// GetDetail 詳細取得
func (ph *projectHandler) GetDetail() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if ok := ph.ru.HasRole(c.GetAuthorID(), constant.GetOperation); !ok {
			return c.HasNotRoleResponse()
		}
		id := ph.getParamID(c)
		project, err := ph.pu.GetDetail(id)
		if err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		data := format.ProjectDetailFormat(project)
		return c.CustomResponse(http.StatusOK, data)
	})
}

// Create 作成
func (ph *projectHandler) Create() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if ok := ph.ru.HasRole(c.GetAuthorID(), constant.CreateOperation); !ok {
			return c.HasNotRoleResponse()
		}
		request := &validation.ProjectRequest{}
		if ok, message := c.BindValidate(request, validation.ProjectMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}
		if err := ph.pu.Create(request); err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"success": "成功",
		})
	})
}

// Update 更新
func (ph *projectHandler) Update() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if ok := ph.ru.HasRole(c.GetAuthorID(), constant.UpdateOperation); !ok {
			return c.HasNotRoleResponse()
		}
		request := &validation.ProjectRequest{}
		if ok, message := c.BindValidate(request, validation.ProjectMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}
		id := ph.getParamID(c)
		if err := ph.pu.Update(id, request); err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"success": "成功",
		})
	})
}

// Delete 削除
func (ph *projectHandler) Delete() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if ok := ph.ru.HasRole(c.GetAuthorID(), constant.DeleteOperation); !ok {
			return c.HasNotRoleResponse()
		}
		id := ph.getParamID(c)
		if err := ph.pu.Delete(id); err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}
		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"success": "成功",
		})
	})
}

// getParamID URLからidを取得する
func (ph *projectHandler) getParamID(c *context.CustomContext) uint32 {
	tmpId, _ := strconv.Atoi(c.Param("id"))
	id := uint32(tmpId)
	return id
}
