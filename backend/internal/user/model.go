package user

type Role struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	RoleID   uint   `json:"role_id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Phone    string `json:"phone"`
}

func (User) TableName() string {
	return "users"
}
