package delegate

import (
	"context"
	"errors"
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent"
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent/user_info"
	"github.com/HaleNing/Ning_JobBoard/src/database"
	"log"
)

func IsExistForUserName(userName string) bool {
	exist, err := database.DBConn.User_info.Query().Where(user_info.UserNameEQ(userName)).Exist(context.Background())
	if err != nil {
		log.Println(err)
		return false
	}
	return exist
}

func CreateNewUserInfo(username string, salt string, hashPass string) *ent.User_info {
	save, err := database.DBConn.User_info.Create().
		SetUserName(username).
		SetSalt(salt).
		SetPasswdHash(hashPass).
		Save(context.Background())
	if err != nil {
		log.Printf("CreateNewUserInfo have failure:[%v]", err)
		return nil
	}
	return save
}

func GetInfoByName(username string) (*ent.User_info, error) {
	exist := IsExistForUserName(username)
	if !exist {
		return nil, errors.New("user is never exist")
	}

	first, err := database.DBConn.User_info.Query().Where(user_info.UserNameEQ(username)).First(context.Background())
	if err != nil {
		log.Printf("getInfoByName have failure:[%v]", err)
		return nil, errors.New("user query failure")

	}
	return first, err
}
