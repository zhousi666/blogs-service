package blogs

import (
	"fmt"

	"blogs-service/common"
)

func SignUp(userName string, password string, sex int64, telNum string) error {
	db := common.DBBegin()
	defer db.DBRollback()
	userInfo := &UserInfo{}
	userInfo.UserName = userName
	userInfo.Password = password
	userInfo.Sex = sex
	userInfo.TelNum = telNum

	err := db.Create(userInfo).Error
	if err != nil {
		common.Logger.Error("creater user error! ")
		return err
	}
	return nil
}

func SignIn(userName string, password string) error {
	db := common.DBBegin()
	defer db.DBRollback()
	notFound := db.Where("username = ?  AND password = ?", userName, password).Find(&UserInfo{}).RecordNotFound()
	if notFound {
		return fmt.Errorf("username or password error! please try angain or sign up.")
	}
	return nil
}
