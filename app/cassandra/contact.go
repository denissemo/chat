package cassandra

import "time"

type Contact struct {
	Username    string    `json:"username"`
	LastSeen    time.Time `json:"last_seen"`
	Name        string    `json:"name"`
	UserContact string    `json:"user_contact"`
}

func FindAllContacts(username string) *[]Contact {
	var contactList []Contact
	temp := map[string]interface{}{}

	q := "SELECT * FROM contacts where username = ? ORDER BY last_seen DESC ALLOW FILTERING"
	contacts := Session.Query(q, username).Iter()
	for contacts.MapScan(temp) {
		contactList = append(contactList, Contact{
			Username:    temp["username"].(string),
			LastSeen:    temp["last_seen"].(time.Time),
			Name:        temp["name"].(string),
			UserContact: temp["user_contact"].(string),
		})
		temp = map[string]interface{}{}
	}

	return &contactList
}
