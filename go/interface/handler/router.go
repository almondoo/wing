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
	InitRoleRouting(RoleHandler)
}

type routing struct {
	e          *echo.Group
	auth       auth.AuthInterface
	token      auth.TokenInterface
	adminToken auth.TokenInterface
}

func NewRouting(e *echo.Group, auth auth.AuthInterface, token auth.TokenInterface) Routing {
	return &routing{e: e, auth: auth, token: token}
}

func (r *routing) InitCommonRouting() {
	r.e.GET("/api/first", context.CastContext(func(c *context.CustomContext) error {
		return c.CustomResponse(http.StatusOK, map[string]string{
			"csrf": c.Get("csrf").(string),
		})
	}))
}

// ユーザー関連処理
func (r *routing) InitAuthUserRouting(userHandler UserHanlder) {
	// グループ
	base := r.e.Group("/api/user")
	// jwtが必要なルート
	jwt := base.Group("", middleware.AuthMiddleware(r.auth, r.token))

	// ログイン処理
	base.POST("/login", userHandler.Login())
	jwt.POST("/logout", userHandler.Logout())

	// ユーザー処理
	base.POST("", userHandler.Register())
	jwt.PUT("", userHandler.Edit())

	// jwt認証確認
	jwt.GET("/check", context.CastContext(func(c *context.CustomContext) error {
		return c.CustomResponse(http.StatusOK, "jwt認証")
	}))

}

func (r *routing) InitRoleRouting(roleHandler RoleHandler) {
	base := r.e.Group("/api/role")
	jwt := base.Group("", middleware.AuthMiddleware(r.auth, r.token))

	jwt.POST("", roleHandler.Create())
	jwt.PUT("", roleHandler.Update())
}
