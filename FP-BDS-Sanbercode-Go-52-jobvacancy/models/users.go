package models

type Users struct {
	Id_User  int    `json:"id_user"`
	Email    string `form:"email" `
	Password string `form:"password"`
	Username string `form:"username"`
	Role     string `form:"role"`
}