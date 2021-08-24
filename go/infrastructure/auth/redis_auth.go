package auth

import (
	"fmt"
	"strconv"
	"time"
	"wing/infrastructure/storedb"
)

type AuthInterface interface {
	CreateAuth(int, *TokenDetails) error
	AuthValid(string) bool
	FetchAuth(string) (uint64, error)
	DeleteRefresh(string) error
	DeleteTokens(*AccessDetails) error
	DeleteRemainingToken(string) error
}

type ClientData struct {
	client *storedb.RedisService
}

var _ AuthInterface = &ClientData{}

func NewRedisAuth(client *storedb.RedisService) AuthInterface {
	return &ClientData{client: client}
}

func NewStoreAuth(client *storedb.DataStoreService) AuthInterface {
	return &DataStoreData{client: client}
}

//Save token metadata to Redis
func (cd *ClientData) CreateAuth(userid int, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	if err := cd.client.Set(td.TokenUuid, strconv.Itoa(userid), at.Sub(now)); err != nil {
		return err
	}
	if err := cd.client.Set(td.RefreshUuid, strconv.Itoa(userid), rt.Sub(now)); err != nil {
		return err
	}

	return nil
}

// チェック
func (cd *ClientData) AuthValid(tokenUuid string) bool {
	if _, err := cd.FetchAuth(tokenUuid); err != nil {
		return false
	}
	return true
}

// jwtデータがあるかチェック
func (cd *ClientData) FetchAuth(tokenUuid string) (uint64, error) {
	userid, err := cd.client.Get(tokenUuid)
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

// リフレッシュトークンを削除
func (cd *ClientData) DeleteRefresh(refreshUuid string) error {
	// refreshトークン削除
	err := cd.client.Delete(refreshUuid)
	if err != nil {
		return err
	}
	return nil
}

// access & refresh 削除
func (cd *ClientData) DeleteTokens(authD *AccessDetails) error {
	// refresh_uuidを作成
	refreshUuid := fmt.Sprintf("%s++%d", authD.TokenUuid, authD.UserId)

	// access_tokenを削除
	if err := cd.client.Delete(authD.TokenUuid); err != nil {
		return err
	}

	// refresh_tokenを削除
	if err := cd.client.Delete(refreshUuid); err != nil {
		return err
	}

	return nil
}

// 残ったトークンを削除
func (cd *ClientData) DeleteRemainingToken(name string) error {
	if err := cd.client.Delete(name); err != nil {
		return err
	}
	return nil
}
