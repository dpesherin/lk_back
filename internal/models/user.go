package models

type User struct {
	ID       int64  `json:"ID" db:"ID"`
	Login    string `json:"Login" db:"LOGIN"`
	Active   bool   `json:"Active" db:"ACTIVE"`
	Email    string `json:"Email" db:"EMAIL"`
	Name     string `json:"Name" db:"NAME"`
	LastName string `json:"Last_Name" db:"LAST_NAME"`
	Avatar   string `json:"Avatar" db:"AVATAR"`
	Admin    bool   `json:"Admin" db:"ADMIN"`
}

type UserData struct {
	ID       int64  `json:"ID" db:"ID"`
	Login    string `json:"Login" db:"LOGIN"`
	Active   bool   `json:"Active" db:"ACTIVE"`
	Email    string `json:"Email" db:"EMAIL"`
	Name     string `json:"Name" db:"NAME"`
	LastName string `json:"Last_Name" db:"LAST_NAME"`
	Avatar   string `json:"Avatar" db:"AVATAR"`
	Admin    bool   `json:"Admin" db:"ADMIN"`
	Password string `json:"Password" db:"PASS"`
}
