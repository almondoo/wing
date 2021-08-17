package handler

import (
	"net/http"
	"wing/infrastructure/auth"
	"wing/interface/context"
	"wing/interface/middleware"

	"github.com/labstack/echo/v4"
)

type Routing interface {
	InitCommonRouting()
	InitAuthUserRouting(UserHanlder)
}

type routing struct {
	e          *echo.Group
	auth       auth.AuthInterface
	userToken  auth.TokenInterface
	adminToken auth.TokenInterface
}

const (
	AuthorTypeUser   = "user"
	AuthorTypeArtist = "artist"
	AuthorTypeAdmin  = "admin"
)

func NewRouting(e *echo.Group, auth auth.AuthInterface, userToken auth.TokenInterface, adminToken auth.TokenInterface) Routing {
	return &routing{e: e, auth: auth, userToken: userToken, adminToken: adminToken}
}

func (r *routing) InitCommonRouting() {
	r.e.GET("/api/first", context.CastContext(func(c *context.CustomContext) error {
		return c.CustomResponse(http.StatusOK, map[string]string{
			"csrf": c.Get("csrf").(string),
		})
	}))
}

//- ユーザー関連処理
func (r *routing) InitAuthUserRouting(userHandler UserHanlder) {
	//- グループ
	base := r.e.Group("/api/user")

	//- ログイン処理
	base.POST("/login", userHandler.Login())

	//- アカウント作成
	base.POST("/register", userHandler.Register())

	//- jwtが必要なルート
	jwt := base.Group("", middleware.AuthMiddleware(r.auth, r.userToken))

	jwt.GET("/check", context.CastContext(func(c *context.CustomContext) error {
		return c.CustomResponse(http.StatusOK, "check")
	}))

	jwt.POST("/logout", userHandler.Logout())

	jwt.POST("/edit", userHandler.Edit())
}
