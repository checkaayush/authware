package model

type User struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Role  string `json:"role,omitempty"`
}
