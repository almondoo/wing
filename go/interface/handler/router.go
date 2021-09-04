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
	// InitRoleRouting(RoleHandler)
	InitTaskPriorityRouting(TaskPriorityHandler)
	InitTaskStatusRouting(TaskStatusHandler)
	InitProjectRouting(ProjectHandler)
	InitTaskRouting(TaskHandler)
	InitTaskChildRouting(TaskChildHandler)
}

type routing struct {
	e     *echo.Group
	auth  auth.AuthInterface
	token auth.TokenInterface
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

	// ログイン
	base.POST("/login", userHandler.Login())
	// ログアウト
	jwt.POST("/logout", userHandler.Logout())

	// ユーザー作成
	base.POST("", userHandler.Register())
	// ユーザー修正
	jwt.PUT("", userHandler.Edit())

	// jwt認証確認
	jwt.GET("/check", context.CastContext(func(c *context.CustomContext) error {
		return c.CustomResponse(http.StatusOK, "jwt認証")
	}))

}

// 現状使わない
// InitRoleRouting Roleのルーティング
// func (r *routing) InitRoleRouting(roleHandler RoleHandler) {
// 	base := r.e.Group("/api/role")
// 	jwt := base.Group("", middleware.AuthMiddleware(r.auth, r.token))

// 	// 全て取得
// 	jwt.GET("", roleHandler.Get())
// 	// 詳細取得
// 	jwt.GET("/:id", roleHandler.GetDetail())
// 	// 登録
// 	jwt.POST("", roleHandler.Create())
// 	// 更新
// 	jwt.PUT("/:id", roleHandler.Update())
// 	// 削除
// 	jwt.DELETE("/:id", roleHandler.Delete())
// }

// InitTaskPriorityRouting ルーティング
func (r *routing) InitTaskPriorityRouting(taskPriorityHandler TaskPriorityHandler) {
	base := r.e.Group("/api/task-priority")
	jwt := base.Group("", middleware.AuthMiddleware(r.auth, r.token))

	// 取得
	jwt.GET("", taskPriorityHandler.Get())
	// 詳細取得
	jwt.GET("/:id", taskPriorityHandler.GetDetail())
	// 作成
	jwt.POST("", taskPriorityHandler.Create())
	// 更新
	jwt.PUT("/:id", taskPriorityHandler.Update())
	// 削除
	jwt.DELETE("/:id", taskPriorityHandler.Delete())
}

// InitTaskStatusRouting ルーティング
func (r *routing) InitTaskStatusRouting(taskStatusHandler TaskStatusHandler) {
	base := r.e.Group("/api/task-status")
	jwt := base.Group("", middleware.AuthMiddleware(r.auth, r.token))

	// 取得
	jwt.GET("", taskStatusHandler.Get())
	// 詳細取得
	jwt.GET("/:id", taskStatusHandler.GetDetail())
	// 作成
	jwt.POST("", taskStatusHandler.Create())
	// 更新
	jwt.PUT("/:id", taskStatusHandler.Update())
	// 削除
	jwt.DELETE("/:id", taskStatusHandler.Delete())
}

// InitProjectRouting ルーティング
func (r *routing) InitProjectRouting(projectHandler ProjectHandler) {
	base := r.e.Group("/api/project")
	jwt := base.Group("", middleware.AuthMiddleware(r.auth, r.token))

	// 取得
	jwt.GET("", projectHandler.Get())
	// 詳細取得
	jwt.GET("/:id", projectHandler.GetDetail())
	// 作成
	jwt.POST("", projectHandler.Create())
	// 更新
	jwt.PUT("/:id", projectHandler.Update())
	// 削除
	jwt.DELETE("/:id", projectHandler.Delete())
}

// InitTaskRouting ルーティング
func (r *routing) InitTaskRouting(taskHandler TaskHandler) {
	base := r.e.Group("/api/task")
	jwt := base.Group("", middleware.AuthMiddleware(r.auth, r.token))

	// 取得
	jwt.GET("", taskHandler.Get())
	// 詳細取得
	jwt.GET("/:id", taskHandler.GetDetail())
	// 作成
	jwt.POST("", taskHandler.Create())
	// 更新
	jwt.PUT("/:id", taskHandler.Update())
	// 削除
	jwt.DELETE("/:id", taskHandler.Delete())
}

// InitTaskChildRouting ルーティング
func (r *routing) InitTaskChildRouting(taskChildHandler TaskChildHandler) {
	base := r.e.Group("/api/task-child")
	jwt := base.Group("", middleware.AuthMiddleware(r.auth, r.token))

	// 取得
	jwt.GET("", taskChildHandler.Get())
	// 詳細取得
	jwt.GET("/:id", taskChildHandler.GetDetail())
	// 作成
	jwt.POST("", taskChildHandler.Create())
	// 更新
	jwt.PUT("/:id", taskChildHandler.Update())
	// 削除
	jwt.DELETE("/:id", taskChildHandler.Delete())
}
