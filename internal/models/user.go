package models

type User struct {
	ID       int64  `json:"id" db:"ID"`
	Login    string `json:"login" db:"LOGIN"`
	Active   bool   `json:"active" db:"ACTIVE"`
	Email    string `json:"email" db:"EMAIL"`
	Name     string `json:"name" db:"NAME"`
	LastName string `json:"lastName" db:"LAST_NAME"`
	Avatar   string `json:"avatar" db:"AVATAR"`
	Admin    bool   `json:"admin" db:"ADMIN"`
}
