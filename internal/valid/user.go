package valid

type CreateUserRequest struct {
	Name  string `form:"name" binding:"min=6,max=100"`
	Email string `form:"email" binding:"min=6" `

	Password string `form:"password" binding:"min=6" `
}

type UserListRequest struct {
	Name   string `form:"name" binding:"max=100"`
	Status int    `form:"status,default=1" binding:"oneof=0 1"`
}

type UpdateUserRequest struct {
	Id     uint64 `form:"id" binding:"min=1"`
	Status int    `form:"status,default=1" binding:"oneof=0 1"`
}
