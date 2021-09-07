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

func (rs *roleService) HasRole(roleId uint32, operation string) bool {
	role, err := rs.roleRepo.FindByID(roleId)
	if err != nil {
		return false
	}
	switch operation {
	case constant.GetOperation:
		return array.IsArray(getVireActor(), role.Name)

	case constant.CreateOperation:
		return array.IsArray(getCreateActor(), role.Name)

	case constant.UpdateOperation:
		return array.IsArray(getUpdateActor(), role.Name)

	case constant.DeleteOperation:
		return array.IsArray(getDeleteActor(), role.Name)

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

func getVireActor() []interface{} {
	return []interface{}{constant.Developer, constant.Administrator, constant.Editor, constant.Viewer}
}

func getCreateActor() []interface{} {
	return []interface{}{constant.Developer, constant.Administrator, constant.Editor}
}

func getUpdateActor() []interface{} {
	return []interface{}{constant.Developer, constant.Administrator, constant.Editor}
}

func getDeleteActor() []interface{} {
	return []interface{}{constant.Developer, constant.Administrator}
}
