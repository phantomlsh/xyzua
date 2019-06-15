package services

import (
	"strconv"
	"xyzua/daos"
	"xyzua/utils"
)

// return new Token if success, error message if failure
func CheckLogin(username string, password string, random string) (success bool, msg string) {
	u := daos.GetUserByKey("username", username)
	if u == nil {
		success = false
		msg = utils.Msg.LoginErr
		return
	}
	hashPassword := utils.HASH(u.Password, random)

	if u.Role == "DISABLED" {
		success = false
		msg = utils.Msg.DisabledErr
		return
	}

	if password == hashPassword {
		newToken := daos.UpdateToken(u.Id, true)
		if newToken != "" {
			success = true
			msg = newToken
		} else {
			success = false
			msg = utils.Msg.UnknownErr
		}
	} else {
		success = false
		msg = utils.Msg.LoginErr
	}
	return
}

// return new Token if success, error message if failure
func CheckToken(token string) (success bool, msg string) {
	u := daos.GetUserByKey("token", token)
	if u == nil {
		success = false
		msg = utils.Msg.TokenErr
		return
	}
	if u.Expire <= utils.TimeStamp(0) {
		success = false
		msg = utils.Msg.TokenErr
		return
	}
	if u.Role == "DISABLED" {
		success = false
		msg = utils.Msg.DisabledErr
		return
	}
	newToken := daos.UpdateToken(u.Id, false)
	if newToken != "" {
		success = true
		msg = newToken
	} else {
		success = true
		msg = utils.Msg.UnknownErr
	}
	return
}

// return token of user if success, error message if failure
func CheckCode(code string, appName string, appSecret string) (success bool, msg string) {
	app := daos.GetAppByKey("appname", appName)
	if app == nil || app.Secret != appSecret {
		success = false
		msg = utils.Msg.AppErr
		return
	}
	v := daos.GetCodeByKeys("code", code, "", "")
	if v == nil {
		success = false
		msg = utils.Msg.CodeErr
		return
	}
	daos.DeleteCode(v.Id)
	userIdString := strconv.Itoa(int(v.Userid))
	u := daos.GetUserByKey("userid", userIdString)
	if u == nil {
		success = false
		msg = utils.Msg.CodeErr
		return
	}
	if u.Role == "DISABLED" {
		success = false
		msg = utils.Msg.DisabledErr
		return
	}
	success = true
	msg = u.Token
	return
}
