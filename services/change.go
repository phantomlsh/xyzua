package services

import (
	"xyzua/daos"
)

// change password of a user
// return success
func ChangePassword(userId int, newPassword string) bool {
	return daos.UpdatePassword(userId, newPassword)
}

// mark
// return success
func Mark(identity string, _type string) bool {
	v := daos.GetMarkByKeys("identity", identity, "type", _type)
	if v != nil { // already marked
		return true
	}
	v = daos.AddMark(_type, identity)
	if v != nil {
		return true
	}
	return false
}
