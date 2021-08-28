package persistence

import (
	"wing/domain/entity"
	"wing/domain/repository"

	"gorm.io/gorm"
)

type projectRepository struct {
	Conn *gorm.DB
}

func NewProjectRepository(conn *gorm.DB) repository.ProjectRepository {
	return &projectRepository{Conn: conn}
}

// IDで取得
func (rr *projectRepository) FindByID(id uint32) (*entity.Project, error) {
	project := &entity.Project{ID: id}

	if err := rr.Conn.First(&project).Error; err != nil {
		return nil, err
	}

	return project, nil
}

func (rr *projectRepository) Finds() ([]*entity.Project, error) {
	var projects []*entity.Project
	if err := rr.Conn.Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

// Create 作成
func (rr *projectRepository) Create(project *entity.Project) (*entity.Project, error) {
	tx := rr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(&project).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return project, tx.Commit().Error
}

// Update 更新
func (rr *projectRepository) Update(project *entity.Project) (*entity.Project, error) {
	tx := rr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Model(&project).Updates(&project).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return project, tx.Commit().Error
}

// Delete 削除
func (rr *projectRepository) Delete(id uint32) error {
	tx := rr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Delete(&entity.Project{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
