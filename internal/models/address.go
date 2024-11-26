package models

type Addr struct {
	ID       int64  `json:"id" db:"ID"`
	Type     string `json:"type" db:"TYPE"`
	Index    string `json:"index" db:"INDEX"`
	Country  string `json:"country" db:"COUNTRY"`
	Province string `json:"province" db:"PROVINCE"`
	City     string `json:"city" db:"CITY"`
	Street   string `json:"street" db:"STREET"`
	House    string `json:"house" db:"HOUSE"`
	Flat     string `json:"flat" db:"FLAT"`
}
