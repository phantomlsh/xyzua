package daos

import (
	"xyzua/models"
)

// GetTicketByKey retrieves Ticket by Key.
// return nil if Ticket doesn't exist
func GetTicketByKey(key string, value string) *models.Ticket {
	var query = make(map[string]string)
	query[key] = value

	l, _ := models.GetAllTicket(query, []string{}, []string{}, []string{}, 0, 0)

	if l != nil {
		u := l[0].(models.Ticket)
		return &u
	} else {
		return nil
	}
}
