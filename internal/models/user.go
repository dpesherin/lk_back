package models

type User struct {
	ID       int64  `json:"ID" db:"ID"`
	Login    string `json:"Login" db:"Login"`
	Active   bool   `json:"Active" db:"Active"`
	Email    string `json:"Email" db:"Email"`
	Name     string `json:"Name" db:"Name"`
	LastName string `json:"Last_Name" db:"Last_Name"`
	Avatar   string `json:"Avatar" db:"Avatar"`
}
