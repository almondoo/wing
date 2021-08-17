package auth

import (
	"fmt"
	"strconv"
	"time"
	"wing/infrastructure/storedb"
)

type DataStoreData struct {
	client *storedb.DataStoreService
}

var _ AuthInterface = &DataStoreData{}

func NewDataStoreAuth(client *storedb.DataStoreService) AuthInterface {
	return &DataStoreData{client: client}
}

//Save token metadata to Redis
func (cd *DataStoreData) CreateAuth(userid int, td *TokenDetails) error {
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

func (cd *DataStoreData) AuthValid(tokenUuid string) bool {
	if _, err := cd.FetchAuth(tokenUuid); err != nil {
		return false
	}
	return true
}

//Check the metadata saved
func (cd *DataStoreData) FetchAuth(tokenUuid string) (uint64, error) {
	userid, err := cd.client.Get(tokenUuid)
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid.(string), 10, 64)
	return userID, nil
}

func (cd *DataStoreData) DeleteRefresh(refreshUuid string) error {
	//delete refresh token
	err := cd.client.Delete(refreshUuid)
	if err != nil {
		return err
	}
	return nil
}

//Once a user row in the token table
func (cd *DataStoreData) DeleteTokens(authD *AccessDetails) error {
	//get the refresh uuid
	refreshUuid := fmt.Sprintf("%s++%d", authD.TokenUuid, authD.UserId)
	//delete access token

	if err := cd.client.Delete(authD.TokenUuid); err != nil {
		return err
	}
	//delete refresh token
	if err := cd.client.Delete(refreshUuid); err != nil {
		return err
	}
	//When the record is deleted, the return value is 1

	return nil
}

func (cd *DataStoreData) DeleteRemainingToken(name string) error {
	if err := cd.client.Delete(name); err != nil {
		return err
	}
	return nil
}
