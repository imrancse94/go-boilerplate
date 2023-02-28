package Services

import "go-boilerplate/models"

func GetPermissionByUserId(userId int) models.Permission {
	return models.GetRolePageByUserId(userId)
}
