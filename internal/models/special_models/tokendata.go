package special_models

type TokenData struct {
	ID    int64  `json:"id"`
	Login string `json:"login"`
	Admin bool   `json:"admin"`
}
