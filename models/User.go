package models

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `gorm:"type:varchar(100);not null" json:"name"`
	Email     string `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password  string `gorm:"type:varchar(100);not null" json:"-"` // hidden field
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetUserByEmail(email string) User {
	var user User
	err := DB.Where("email = ?", email).First(&user).Error
	if err != nil {

	}
	return user
}

func GetUserById(id int) User {
	var user User
	err := DB.First(&user, id).Error
	if err != nil {
		panic(err)
	}
	return user
}
