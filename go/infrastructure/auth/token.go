package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// AccesssExpires はアクセストークンの制限時間
const AccessExpires = time.Minute * 15

// RefreshExpires はアクセストークンの制限時間
const RefreshExpires = time.Hour * 24 * 30

type (
	Token struct{}

	TokenInterface interface {
		CreateToken(int) (*TokenDetails, error)
		ExtractTokenMetadata(string) (*AccessDetails, error)
	}

	// 暗号化されたjwtトークン
	TokenData struct {
		AccessToken  string
		RefreshToken string
	}

	// アクセストークンのuuid
	AccessDetails struct {
		TokenUuid string
		UserId    uint64
	}

	TokenDetails struct {
		AccessToken  string
		RefreshToken string
		TokenUuid    string
		RefreshUuid  string
		AtExpires    int64
		RtExpires    int64
	}
)

// 構造体を生成する
func NewToken() *Token {
	return &Token{}
}

// トークンのインターファイスを実装する
var _ TokenInterface = &Token{}

func (t *Token) CreateToken(userId int) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(AccessExpires).Unix()
	td.TokenUuid = uuid.New().String()

	td.RtExpires = time.Now().Add(RefreshExpires).Unix()
	td.RefreshUuid = td.TokenUuid + "++" + strconv.Itoa(userId)

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["author_type"] = "user"
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.TokenUuid
	atClaims["author_id"] = userId
	atClaims["exp"] = td.AtExpires
	atClaims["iat"] = time.Now()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["author_type"] = "user"
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["author_id"] = userId
	rtClaims["exp"] = td.RtExpires
	atClaims["iat"] = time.Now()
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

// トークンデータ出す
func (t *Token) ExtractTokenMetadata(tokenString string) (*AccessDetails, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, errors.New("アサーションに失敗しました。")
		}
		id, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["author_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			TokenUuid: accessUuid,
			UserId:    id,
		}, nil
	}
	return nil, err
}

/* Access Token start */
// トークンを正しいかチェック
func TokenValid(tokenString string) bool {
	if _, err := VerifyToken(tokenString); err != nil {
		return false
	}
	return true
}

// jwtデータ取得
func FetchTokenClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "トークンデータ取得エラー")
	}
	return claims, nil
}

// トークンが正しいか
func VerifyToken(tokenString string) (*jwt.Token, error) {
	// 検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("検証エラー")
	}
	return token, nil
}

/* Access Token end */

/* Refresh Token start */
// トークンが正しいかチェック
func RefreshTokenValid(refreshTokenString string) bool {
	if _, err := VerifyRefreshToken(refreshTokenString); err != nil {
		return false
	}
	return true
}

// tokenデータ取得
func FetchRefreshTokenClaims(refreshTokenString string) (jwt.MapClaims, error) {
	token, err := VerifyRefreshToken(refreshTokenString)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "トークンデータ取得エラー")
	}
	return claims, nil
}

func VerifyRefreshToken(refreshTokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

/* Refresh Token end */
