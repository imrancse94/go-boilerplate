package models

type RolePage struct {
	ID                int    `gorm:"primary_key" json:"id"`
	UserID            int    `gorm:"type:int(11);not null" json:"user_id"`
	PermissionVersion int    `gorm:"type:int(11);not null" json:"permission_version"`
	PageID            int    `gorm:"type:int(11);not null" json:"page_id"`
	PageName          string `gorm:"type:varchar;not null" json:"page_name"`
	ParentID          int    `gorm:"type:int(11);not null" json:"parent_id"`
	IsIndex           int    `gorm:"type:int(11);not null" json:"is_index"`
	Icon              string `gorm:"type:varchar;not null" json:"icon"`
	PermissionName    string `gorm:"type:varchar;not null" json:"permission_name"`
	Submenu           map[int]interface{}
}

type Permission struct {
	PermissionList  []string            `json:"permission_list"`
	PermissionAssoc map[int]interface{} `json:"permission_assoc"`
}

func GetRolePageByUserId(id int) Permission {
	var rolePage []RolePage
	sql := `SELECT DISTINCT menu_submenu_permissions.id,
                        users.id AS user_id,
                        users.permission_version,
                        menu_submenu_permissions.id,
                        menu_submenu_permissions.name AS page_name,
                        menu_submenu_permissions.parent_id,
                        menu_submenu_permissions.is_index,
                        menu_submenu_permissions.icon,
                        menu_submenu_permissions.permission_name FROM users
                        INNER JOIN user_usergroups ON users.id = user_usergroups.user_id
                        INNER JOIN usergroups ON user_usergroups.usergroup_id = usergroups.id
                        INNER JOIN usergroup_roles ON usergroups.id = usergroup_roles.usergroup_id
                        INNER JOIN roles ON usergroup_roles.role_id = roles.id
                        INNER JOIN role_menu_submenu_permissions ON roles.id = role_menu_submenu_permissions.role_id
                        INNER JOIN menu_submenu_permissions ON role_menu_submenu_permissions.menu_submenu_permission_id = menu_submenu_permissions.id
                        WHERE users.id = ?`

	err := DB.Raw(sql, id).Scan(&rolePage).Error

	if err != nil {
		panic(err)
	}
	Permission := Permission{}
	indexList := make(map[int]interface{})
	for rolePageIndex := range rolePage {
		if rolePage[rolePageIndex].PermissionName != "" {
			Permission.PermissionList = append(Permission.PermissionList, rolePage[rolePageIndex].PermissionName)
		}

		if rolePage[rolePageIndex].IsIndex == 1 {
			indexList[rolePage[rolePageIndex].ParentID] = rolePage[rolePageIndex]
		}
	}

	Permission.PermissionAssoc = makeSideBar(rolePage, 0, indexList)

	return Permission
}

func makeSideBar(elements []RolePage, parentId int, indexList map[int]interface{}) map[int]interface{} {
	var branch = make(map[int]interface{})
	//fmt.Println("elements", branch)
	for elementIndex := range elements {
		// array key not exist do not proceed
		if !Isset(elements, elementIndex) {
			continue
		}

		if elements[elementIndex].ParentID == parentId {

			children := makeSideBar(elements, elements[elementIndex].ID, indexList)

			if elements[elementIndex].ID > 0 && indexList[elements[elementIndex].ID] != nil {
				elements[elementIndex].PermissionName = indexList[elements[elementIndex].ID].(RolePage).PermissionName
				elements[elementIndex].IsIndex = 1
			}

			if children != nil {
				elements[elementIndex].Submenu = make(map[int]interface{})
				elements[elementIndex].Submenu = children
			}

			branch[elements[elementIndex].ID] = elements[elementIndex]
			elements = RemoveIndex(elements, elementIndex)
		}
	}

	return branch
}

func RemoveIndex(s []RolePage, index int) []RolePage {
	return append(s[:index], s[index+1:]...)
}

func Isset(arr []RolePage, index int) bool {
	return len(arr) > index
}
