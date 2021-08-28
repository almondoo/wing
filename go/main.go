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
	if os.Getenv("ENV") == "local" {
		repositories.Seeder()
	}

	// Newトークン
	token := auth.NewToken()
	group := e.Group("")
	routing := handler.NewRouting(group, newAuth, token)

	// 共通Routing
	routing.InitCommonRouting()

	// User関連のRouting
	userService := service.NewUserService(repositories.User, newAuth, token)
	userUsecase := usecase.NewUserUsecase(userService, newAuth, token)
	userHandler := handler.NewUserHandler(userUsecase)
	routing.InitAuthUserRouting(userHandler)

	// Role関連のRouting
	roleService := service.NewRoleService(repositories.Role)
	roleUsecase := usecase.NewRoleUsecase(roleService)
	roleHandler := handler.NewRoleHandler(roleUsecase)
	routing.InitRoleRouting(roleHandler)

	// Role関連のRouting
	taskPriorityService := service.NewTaskPriorityService(repositories.TaskPriority)
	taskPriorityUsecase := usecase.NewTaskPriorityUsecase(taskPriorityService)
	taskPriorityHandler := handler.NewTaskPriorityHandler(taskPriorityUsecase)
	routing.InitTaskPriorityRouting(taskPriorityHandler)

	// Role関連のRouting
	taskStatusService := service.NewTaskStatusService(repositories.TaskStatus)
	taskStatusUsecase := usecase.NewTaskStatusUsecase(taskStatusService)
	taskStatusHandler := handler.NewTaskStatusHandler(taskStatusUsecase)
	routing.InitTaskStatusRouting(taskStatusHandler)

	// Role関連のRouting
	projectService := service.NewProjectService(repositories.Project)
	projectUsecase := usecase.NewProjectUsecase(projectService)
	projectHandler := handler.NewProjectHandler(projectUsecase)
	routing.InitProjectRouting(projectHandler)

	// Role関連のRouting
	taskService := service.NewTaskService(repositories.Task)
	taskUsecase := usecase.NewTaskUsecase(taskService)
	taskHandler := handler.NewTaskHandler(taskUsecase)
	routing.InitTaskRouting(taskHandler)

	// Role関連のRouting
	taskChildService := service.NewTaskChildService(repositories.TaskChild)
	taskChildUsecase := usecase.NewTaskChildUsecase(taskChildService)
	taskChildHandler := handler.NewTaskChildHandler(taskChildUsecase)
	routing.InitTaskChildRouting(taskChildHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
