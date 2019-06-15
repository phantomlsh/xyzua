package daos

import (
	"xyzua/models"
	"xyzua/utils"
)

// GetUserByKey retrieves User by Key.
// return nil if User doesn't exist
func GetUserByKey(key string, value string) *models.User {
	var query = make(map[string]string)
	query[key] = value

	l, _ := models.GetAllUser(query, []string{}, []string{}, []string{}, 0, 0)

	if l != nil {
		u := l[0].(models.User)
		return &u
	} else {
		return nil
	}
}

// UpdateToken by userid
// @Success new token
// @Failure empty string
func UpdateToken(userId int, updateExpire bool) string {
	v, err := models.GetUserById(userId)
	if err != nil {
		return ""
	}
	newToken := utils.RandomString(32)
	v.Token = newToken
	if updateExpire == true {
		v.Expire = utils.TimeStamp(3600)
	}
	err = models.UpdateUserById(v)
	if err != nil {
		return ""
	}
	return newToken
}

// UpdatePassword by userid
// return success
func UpdatePassword(userId int, newPassword string) bool {
	v, err := models.GetUserById(userId)
	if err != nil {
		return false
	}
	v.Password = newPassword
	err = models.UpdateUserById(v)
	if err != nil {
		return false
	}
	return true
}
