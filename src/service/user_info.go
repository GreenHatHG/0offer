package service

import (
	"0offer/global"
	"0offer/model"
)

func SelectAllUserInfo() ([]model.UserInfo, error) {
	var userInfo []model.UserInfo
	err := global.DB.Find(&userInfo).Error
	return userInfo, err
}

func BatchInsertUserInfo(userInfos []model.UserInfo) ([]model.UserInfo, error) {
	if len(userInfos) != 0 {
		err := global.DB.Create(&userInfos).Error
		return userInfos, err
	}
	return nil, nil
}
