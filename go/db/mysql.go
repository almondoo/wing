package db

import (
	"log"
	"os"
	"time"
	"wing/db/seed"
	"wing/domain/entity"
	"wing/domain/repository"
	"wing/infrastructure/persistence"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repositories struct {
	User repository.UserRepository
	Role repository.RoleRepository
	DB   *gorm.DB
}

//- connetion確認func
func InitDB() (*Repositories, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      false,
		},
	)

	dns := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@" + os.Getenv("DB_PROTO") + "/" + os.Getenv("DB") + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	//- 自動的にテーブルを生成する
	db.AutoMigrate(&entity.Role{}, &entity.User{})

	return &Repositories{
		User: persistence.NewUserRepository(db),
		Role: persistence.NewRoleRepository(db),
		DB:   db,
	}, nil

}

func (r *Repositories) Close() error {
	db, err := r.DB.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func (r *Repositories) Seeder() {
	db := r.DB
	seed.NewRoleSeeder(db).Seeder()
	seed.NewUserSeeder(db).Seeder()
}