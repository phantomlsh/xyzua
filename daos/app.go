package daos

import (
	"xyzua/models"
)

// GetAppByKey retrieves App by Key.
// return nil if App doesn't exist
func GetAppByKey(key string, value string) *models.App {
	var query = make(map[string]string)
	query[key] = value

	l, _ := models.GetAllApp(query, []string{}, []string{}, []string{}, 0, 0)

	if l != nil {
		u := l[0].(models.App)
		return &u
	} else {
		return nil
	}
}

// GetAppList retrieves all Apps
func GetAppList() (ml []models.App) {
	l, _ := models.GetAllApp(map[string]string{}, []string{}, []string{}, []string{}, 0, 0)

	for _, v := range l {
		u := v.(models.App)
		u.Secret = "*"   // hide secret from list
		u.Callback = "*" // hide callback from list
		ml = append(ml, u)
	}
	return
}
