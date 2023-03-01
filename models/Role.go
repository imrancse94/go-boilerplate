package models

type Role struct {
	ID     int    `gorm:"primary_key" json:"id"`
	Title  string `gorm:"type:varchar(100);not null" json:"title"`
	Status int    `gorm:"type:tinyint(4);not null" json:"status"`
}

func GetRoles() []Role {
	var roles []Role
	err := DB.Find(&roles).Error

	if err != nil {
		return nil
	}
	return roles
}

func AddRole(role Role) interface{} {
	err := DB.Create(&role)

	if err != nil {
		return nil
	}
	return role
}
