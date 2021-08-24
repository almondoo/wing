package persistence

import (
	"wing/domain/repository"

	"gorm.io/gorm"
)

type sampleRepository struct {
	Conn *gorm.DB
}

func NewSampleRepository(conn *gorm.DB) repository.SampleRepository {
	return &sampleRepository{Conn: conn}
}

// 作成
func (rr *sampleRepository) Create() error {
	// Createロジック
	return nil
}

// 更新
func (rr *sampleRepository) Update() error {
	// Updateロジック
	return nil
}

// 削除
func (rr *sampleRepository) Delete() error {
	// Deleteロジック
	return nil
}
