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

	//- CustomValidation
	e.Validator = validation.NewValidator()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(middleware.CSRFMiddleware())
	e.Use(middleware.CORSMiddleware())

	// authClient, err := store.NewDataStoreDB(os.Getenv("PROJECT_ID"))
	authClient := storedb.NewRedisDB()
	newAuth := auth.NewRedisAuth(authClient)

	//- Repository
	repositories, err := db.InitDB()
	if err != nil {
		panic(err)
	}
	defer repositories.Close()
	if os.Getenv("ENV") == "local" {
		repositories.Seeder()
	}

	//- validator
	dbValidator := validation.NewValidatorWithDB(repositories.DB)
	e.Use(middleware.CustomContextMiddleware(dbValidator))

	//- Newトークン
	userToken := auth.NewUserToken()
	adminToken := auth.NewAdminToken()
	group := e.Group("")
	routing := handler.NewRouting(group, newAuth, userToken, adminToken)

	//- 共通Routing
	routing.InitCommonRouting()

	//- User関連のRouting
	userService := service.NewUserService(repositories.User, newAuth, userToken)
	userUsecase := usecase.NewUserUsecase(userService, newAuth, userToken)
	userHandler := handler.NewUserHandler(userUsecase)
	routing.InitAuthUserRouting(userHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
