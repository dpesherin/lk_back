package models

type Legal struct {
	ID            int64  `json:"id" db:"ID"`
	Inn           int64  `json:"inn" db:"INN"`
	Kpp           int64  `json:"kpp" db:"KPP"`
	Short         string `json:"short" db:"SHORT"`
	Official      string `json:"official" db:"OFFICIAL"`
	JuridicalAddr *Addr
	ActualAddr    *Addr
	DeliveryAddr  *Addr
}
