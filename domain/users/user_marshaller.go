package users

import "encoding/json"

type PublicUser struct {
	ID          int64  `json:"id"`
	DateCreated string `json:"created_at"`
	Status      string `json:"status"`
}

type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"created_at"`
	Status      string `json:"status"`
}

func (u Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(u))
	for i, user := range u {
		result[i] = user.Marshall(isPublic)
	}
	return result
}

func (u *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:          u.ID,
			DateCreated: u.DateCreated,
			Status:      u.Status,
		}
	}
	usrJSON, _ := json.Marshal(u)
	var privateUser PrivateUser
	json.Unmarshal(usrJSON, &privateUser)
	return privateUser
}
