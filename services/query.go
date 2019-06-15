package services

import (
	"xyzua/daos"
	"xyzua/utils"

	"strings"
)

type Identity struct {
	IsUser   bool
	Role     string
	Identity string
	Id       int
	Raw      interface{}
}

type Callback struct {
	Token string
	Url   string
}

// return identity report with modified information
// return nil if fail
func QueryIdentity(t string, updateToken bool) *Identity {
	var res Identity
	u := daos.GetUserByKey("token", t)
	if u != nil && u.Expire > utils.TimeStamp(0) {
		res.IsUser = true
		res.Role = u.Role
		res.Identity = u.Username
		res.Id = u.Id
		if updateToken == true {
			u.Token = daos.UpdateToken(u.Id, false)
		}
		u.Password = "*" // hide password
		res.Raw = *u
		return &res
	}
	v := daos.GetTicketByKey("ticket", t)
	if v != nil && v.Expire > utils.TimeStamp(0) {
		res.IsUser = false
		res.Role = "TICKET"
		res.Identity = v.Ticket
		res.Id = v.Id
		res.Raw = *v
		return &res
	}
	return nil
}

// Create and return Callback report
// return nil if fail
func QueryCode(userId int, appName string) *Callback {
	app := daos.GetAppByKey("appname", appName)
	if app == nil {
		return nil
	}
	v := daos.AddCode(userId, app.Id)
	if v != nil {
		return &Callback{Url: strings.Replace(app.Callback, "#code#", v.Code, 1)}
	}
	return nil
}

// return App list
func QueryApps() (ml []interface{}) {
	l := daos.GetAppList()
	for _, v := range l {
		ml = append(ml, v)
	}
	return
}

// return ticket information
// return nil if ticket does not exist
func QueryTicket(ticket string) interface{} {
	v := daos.GetTicketByKey("ticket", ticket)
	return v
}

// return mark information
// return empty if mark does not exist
func QueryMarks(identity string) (ml []interface{}) {
	l := daos.GetMarksByKey("identity", identity)
	for _, v := range l {
		ml = append(ml, v)
	}
	return
}
