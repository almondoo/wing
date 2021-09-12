package service

import (
	"wing/domain/entity"
	"wing/domain/repository"
	"wing/interface/validation"
	"wing/pkg/array"
	"wing/utils/constant"
)

type RoleService interface {
	HasRole(roleId uint32, operation string) bool
	IsAdmin(roleId uint32) bool
	Get() ([]*entity.Role, error)
	GetDetail(id uint32) (*entity.Role, error)
	Create(*validation.RoleRequest) error
	Update(uint32, *validation.RoleRequest) error
	Delete(uint32) error
}

type roleService struct {
	roleRepo repository.RoleRepository
}

func NewRoleService(roleRepo repository.RoleRepository) RoleService {
	return &roleService{roleRepo: roleRepo}
}

// HasRole メソッドの権限を持っているか
func (rs *roleService) HasRole(roleId uint32, operation string) bool {
	role, err := rs.roleRepo.FindByID(roleId)
	if err != nil {
		return false
	}
	switch operation {
	case constant.GetOperation:
		return array.IsArray(getViewRole(), role.Name)

	case constant.CreateOperation:
		return array.IsArray(getCreateAndUpdateRole(), role.Name)

	case constant.UpdateOperation:
		return array.IsArray(getCreateAndUpdateRole(), role.Name)

	case constant.DeleteOperation:
		return array.IsArray(getDeleteRole(), role.Name)

	}
	return false
}

// IsAdmin 管理者以上の権限確認
func (rs *roleService) IsAdmin(roleId uint32) bool {
	role, err := rs.roleRepo.FindByID(roleId)
	if err != nil {
		return false
	}
	if array.IsArray(getAdminRole(), role.Name) {
		return true
	}
	return false
}

func (rs *roleService) Get() ([]*entity.Role, error) {
	return rs.roleRepo.Finds()
}
func (rs *roleService) GetDetail(id uint32) (*entity.Role, error) {
	return rs.roleRepo.FindByID(id)
}

func (rs *roleService) Create(request *validation.RoleRequest) error {
	role := &entity.Role{
		Name: request.Name,
	}
	_, err := rs.roleRepo.Create(role)
	return err
}

func (rs *roleService) Update(id uint32, request *validation.RoleRequest) (err error) {
	var role *entity.Role
	role, err = rs.roleRepo.FindByID(id)
	if err != nil {
		return
	}
	_, err = rs.roleRepo.Update(role)
	return
}

func (rs *roleService) Delete(id uint32) (err error) {
	return rs.roleRepo.Delete(id)
}

// getViewRole 閲覧権限を持っているロール
func getViewRole() []interface{} {
	return []interface{}{constant.Developer, constant.Administrator, constant.Editor, constant.Viewer}
}

// getCreateAndUpdateRole 作成・更新権限を持っているロール
func getCreateAndUpdateRole() []interface{} {
	return []interface{}{constant.Developer, constant.Administrator, constant.Editor}
}

// getDeleteRole 削除権限を持っているロール
func getDeleteRole() []interface{} {
	return []interface{}{constant.Developer, constant.Administrator}
}

// getViewRole 管理者権限以上を持っているロール
func getAdminRole() []interface{} {
	return []interface{}{constant.Developer, constant.Administrator}
}
