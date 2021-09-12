package main

import (
	"fmt"
	"os"
	"wing/application/usecase"
	"wing/db"
	"wing/domain/service"
	"wing/infrastructure/auth"
	"wing/infrastructure/storedb"
	"wing/interface/handler"
	"wing/interface/middleware"
	"wing/interface/validation"

	env "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
)

func main() {
	err := env.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		panic(err)
	}

	e := echo.New()

	// CustomValidation
	e.Validator = validation.NewValidator()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(middleware.CSRFMiddleware())
	e.Use(middleware.CORSMiddleware())
	e.Use(middleware.CustomContextMiddleware())

	// authClient, err := store.NewDataStoreDB(os.Getenv("PROJECT_ID"))
	authClient := storedb.NewRedisDB()
	newAuth := auth.NewRedisAuth(authClient)

	// Repository
	repositories, err := db.InitDB()
	if err != nil {
		panic(err)
	}
	defer repositories.Close()
	if err := repositories.Migrations(); err != nil {
		panic(err)
	}
	if os.Getenv("ENV") == "local" {
		repositories.Seeder()
	}

	// Newトークン
	token := auth.NewToken()
	group := e.Group("")
	routing := handler.NewRouting(group, newAuth, token)

	// service作成
	userService := service.NewUserService(repositories.User, newAuth, token)
	roleService := service.NewRoleService(repositories.Role)
	taskPriorityService := service.NewTaskPriorityService(repositories.TaskPriority)
	taskStatusService := service.NewTaskStatusService(repositories.TaskStatus)
	projectService := service.NewProjectService(repositories.Project)
	taskService := service.NewTaskService(repositories.Task)
	taskChildService := service.NewTaskChildService(repositories.TaskChild)

	// usecase作成
	userUsecase := usecase.NewUserUsecase(userService, roleService, newAuth, token)
	roleUsecase := usecase.NewRoleUsecase(roleService, userService)
	taskPriorityUsecase := usecase.NewTaskPriorityUsecase(taskPriorityService)
	taskStatusUsecase := usecase.NewTaskStatusUsecase(taskStatusService)
	projectUsecase := usecase.NewProjectUsecase(projectService)
	taskUsecase := usecase.NewTaskUsecase(taskService)
	taskChildUsecase := usecase.NewTaskChildUsecase(taskChildService)

	// 現状使わないかも
	// Role関連のRouting
	// roleHandler := handler.NewRoleHandler(roleUsecase, roleUsecase)
	// routing.InitRoleRouting(roleHandler)

	// 共通Routing
	routing.InitCommonRouting()

	// User関連のRouting
	userHandler := handler.NewUserHandler(userUsecase, roleUsecase)
	routing.InitAuthUserRouting(userHandler)

	// Role関連のRouting
	taskPriorityHandler := handler.NewTaskPriorityHandler(taskPriorityUsecase, roleUsecase)
	routing.InitTaskPriorityRouting(taskPriorityHandler)

	// Role関連のRouting
	taskStatusHandler := handler.NewTaskStatusHandler(taskStatusUsecase, roleUsecase)
	routing.InitTaskStatusRouting(taskStatusHandler)

	// Role関連のRouting
	projectHandler := handler.NewProjectHandler(projectUsecase, roleUsecase)
	routing.InitProjectRouting(projectHandler)

	// Role関連のRouting
	taskHandler := handler.NewTaskHandler(taskUsecase, roleUsecase)
	routing.InitTaskRouting(taskHandler)

	// Role関連のRouting
	taskChildHandler := handler.NewTaskChildHandler(taskChildUsecase, roleUsecase)
	routing.InitTaskChildRouting(taskChildHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
