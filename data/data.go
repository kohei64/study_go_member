package data

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Belongs  string `json:"belongs"`
	Skills   string `json:"skills"`
}