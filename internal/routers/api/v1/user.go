package v1

// import (
// 	"github.com/gin-gonic/gin"

// 	"github.com/itachilee/gin/pkg/app"
// 	"github.com/itachilee/gin/pkg/convert"
// 	"github.com/itachilee/gin/pkg/errcode"
// )

// type User struct {
// 	Name     string `json:"name" `
// 	Email    string `json:"email" `
// 	Status   int    `json:"status" `
// 	Password string `json:"password" `
// }

// func NewUser() User {
// 	return User{}
// }

// func (a User) Get(c *gin.Context) {}
// func (a User) List(c *gin.Context) {
// 	param := service.UserListRequest{}
// 	response := app.NewResponse(c)
// 	valid, errs := app.BindAndValid(c, &param)
// 	if !valid {
// 		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
// 		return
// 	}
// 	svc := service.New(c.Request.Context())
// 	users, err := svc.GetUserList(&param, &app.Pager{})
// 	if err != nil {

// 		response.ToErrorResponse(errcode.ErrorCreateTagFail.WithDetails(err.Error()))
// 		return
// 	}

// 	response.ToResponse(gin.H{"data": users})
// }
// func (a User) Create(c *gin.Context) {
// 	param := service.CreateUserRequest{}
// 	response := app.NewResponse(c)
// 	valid, errs := app.BindAndValid(c, &param)
// 	if !valid {
// 		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
// 		return

// 	}

// 	svc := service.New(c.Request.Context())
// 	err := svc.CreateUser(&param)
// 	if err != nil {

// 		response.ToErrorResponse(errcode.ErrorCreateTagFail.WithDetails(err.Error()))
// 		return
// 	}

// 	response.ToResponse(gin.H{"code": errcode.Success, "msg": "操作成功"})

// }
// func (a User) Update(c *gin.Context) {

// 	id := c.Param("id")
// 	param := service.UpdateUserRequest{
// 		Id: convert.StrTo(id).MustUInt64(),
// 	}
// 	response := app.NewResponse(c)
// 	valid, errs := app.BindAndValid(c, &param)
// 	if !valid {
// 		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
// 		return

// 	}

// 	svc := service.New(c.Request.Context())
// 	err := svc.UpdateUser(&param)
// 	if err != nil {

// 		response.ToErrorResponse(errcode.ErrorCreateTagFail.WithDetails(err.Error()))
// 		return
// 	}

// 	response.ToResponse(gin.H{"code": errcode.Success, "msg": "操作成功"})

// }
// func (a User) Delete(c *gin.Context) {}
