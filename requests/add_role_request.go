package requests

type AddRoleRequest struct {
	Title  string `json:"title" valid:"required~Title is required"`
	Status string `json:"status" valid:"required~Status is required,int~Status must be integer"`
}
