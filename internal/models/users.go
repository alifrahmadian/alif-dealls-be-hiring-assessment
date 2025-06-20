package models

type User struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	BaseSalary uint64 `json:"base_salary"`
	RoleID int64 `json:"role_id"`
	Role Role `json:"role"`
}