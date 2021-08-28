package middleware

import (
	"fmt"
	"net/http"
	"os"
	"wing/infrastructure/auth"
	"wing/interface/context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CORSMiddleware Next.jsをホスティングしているの場所にアクセスできるようにCORSを設定
func CORSMiddleware() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("CORS_URL")},
		AllowHeaders: []string{
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderXCSRFToken,
			echo.HeaderXRequestedWith,
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
		},
		AllowCredentials: true,
	})
}

// csrf
func CSRFMiddleware() echo.MiddlewareFunc {
	return middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:  "header:X-CSRF-TOKEN",
		CookieDomain: os.Getenv("CORS_DOMAIN"),
		CookieName:   "csrf",
		CookiePath:   "/",
		// CookieSecure:   true, // SSLが常時した時にコメントアウトを取る
		CookieHTTPOnly: true,
	})
}

// AuthMiddleware ユーザー側の認証
func AuthMiddleware(authenticate auth.AuthInterface, token auth.TokenInterface) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return context.CastContext(func(c *context.CustomContext) error {
			fmt.Println("jwt")
			c.SetAuthorID(0)

			// トークンを取得
			tokens := c.GetCookieToken()
			// トークンを確認
			if ok := auth.TokenValid(tokens.AccessToken); ok {
				// エラーではなかったら
				if claims, err := auth.FetchTokenClaims(tokens.AccessToken); err == nil {
					// エラーではなかったら
					if ok := authenticate.AuthValid(claims["access_uuid"].(string)); ok {
						intID := int(claims["author_id"].(float64))
						c.SetAuthorID(uint(intID))
						return next(c)
					}
				}
			}

			// リフレッシュトークンを確認
			if ok := auth.RefreshTokenValid(tokens.RefreshToken); !ok {
				fmt.Println("リフレッシュトークンは正しくない")
				return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
					"data": map[string]string{
						"error": "認証に失敗しました。",
					},
					"status": "ng",
				})
			}
			claims, err := auth.FetchRefreshTokenClaims(tokens.RefreshToken)
			if err != nil {
				fmt.Println("リフレッシュトークンのデータ取得失敗")
				return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
					"data": map[string]string{
						"error": "認証に失敗しました。",
					},
					"status": "ng",
				})
			}
			if ok := authenticate.AuthValid(claims["refresh_uuid"].(string)); !ok {
				fmt.Println("使用できないリフレッシュトークン")
				return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
					"data": map[string]string{
						"error": "認証に失敗しました。",
					},
					"status": "ng",
				})
			}
			authorId := claims["author_id"].(float64)
			cToken, err := token.CreateToken(int(authorId))
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
					"data": map[string]string{
						"error": "トークン生成エラー",
					},
					"status": "ng",
				})
			}
			if err := authenticate.CreateAuth(int(authorId), cToken); err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
					"data": map[string]string{
						"error": "トークンの保存に失敗しました。",
					},
					"status": "ng",
				})
			}
			authenticate.DeleteRemainingToken(claims["refresh_uuid"].(string))
			intID := int(authorId)
			c.SetAuthorID(uint(intID))
			c.SetCookieToken(cToken.AccessToken, cToken.RefreshToken)

			return next(c)
		})
	}
}

// CustomContextMiddleware はカスタマイズされたコンテキストを生成する関数
func CustomContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &context.CustomContext{Context: c}
			return next(cc)
		}
	}
}
