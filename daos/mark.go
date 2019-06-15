package daos

import (
	"xyzua/models"
)

// GetMarksByKey retrieves Marks by Key.
// return empty if Mark doesn't exist
func GetMarksByKey(key string, value string) (ml []models.Mark) {
	var query = make(map[string]string)
	query[key] = value

	l, _ := models.GetAllMark(query, []string{}, []string{}, []string{}, 0, 0)

	for _, v := range l {
		u := v.(models.Mark)
		ml = append(ml, u)
	}
	return
}

// GetMarkByKeys retrieves Mark by Keys.
// return nil if Mark doesn't exist
func GetMarkByKeys(k1 string, v1 string, k2 string, v2 string) *models.Mark {
	var query = make(map[string]string)
	query[k1] = v1
	query[k2] = v2

	l, _ := models.GetAllMark(query, []string{}, []string{}, []string{}, 0, 0)

	if l != nil {
		u := l[0].(models.Mark)
		return &u
	} else {
		return nil
	}
}

// Delete by markid
// return success
func DeleteMark(markId int) bool {
	if err := models.DeleteMark(markId); err == nil {
		return true
	} else {
		return false
	}
}

// UpdatePassword by userid
// return *models.Mark added if success
// return nil if fail
func AddMark(_type string, identity string) *models.Mark {
	var v models.Mark
	v.Type = _type
	v.Identity = identity
	if _, err := models.AddMark(&v); err == nil {
		return &v
	}
	return nil
}
