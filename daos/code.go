package daos

import (
	"strconv"
	"xyzua/models"
	"xyzua/utils"
)

// GetCodeByKey retrieves Code by Keys.
// return nil if Code doesn't exist
func GetCodeByKeys(k1 string, v1 string, k2 string, v2 string) *models.Code {
	var query = make(map[string]string)
	query[k1] = v1
	query[k2] = v2

	l, _ := models.GetAllCode(query, []string{}, []string{}, []string{}, 0, 0)

	if l != nil {
		u := l[0].(models.Code)
		return &u
	} else {
		return nil
	}
}

// Delete by codeid
// return success
func DeleteCode(codeId int) bool {
	if err := models.DeleteCode(codeId); err == nil {
		return true
	} else {
		return false
	}
}

// UpdatePassword by userid
// return *models.Code added if success
// return nil if fail
func AddCode(userId int, appId int) *models.Code {
	userIdString := strconv.Itoa(userId)
	appIdString := strconv.Itoa(appId)
	v := GetCodeByKeys("userid", userIdString, "appid", appIdString)
	var u models.Code
	var err error
	if v != nil {
		u = *v
	} else {
		u.Userid = uint(userId)
		u.Appid = uint(appId)
		_, err = models.AddCode(&u)
		if err != nil {
			return nil
		}
	}
	u.Code = utils.RandomString(64)
	u.Expire = utils.TimeStamp(600)
	err = models.UpdateCodeById(&u)
	if err == nil {
		return &u
	}
	return nil
}
