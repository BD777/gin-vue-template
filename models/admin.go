package models

type Admin struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func GetAdmin(username, password string) (admin *Admin) {
	admin = &Admin{}
	db.Select("id").Where(Admin{Username: username, Password: password}).First(admin)
	if admin.ID > 0 {
		return admin
	}

	return nil
}
