package models

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Obj     interface{} `json:"obj"`
}
