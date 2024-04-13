package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/miniblog/internal/miniblog/biz"
	"github.com/marmotedu/miniblog/internal/miniblog/store"
	"github.com/marmotedu/miniblog/internal/pkg/core"
	"github.com/marmotedu/miniblog/internal/pkg/errno"
	"github.com/marmotedu/miniblog/internal/pkg/log"
	v1 "github.com/marmotedu/miniblog/pkg/api/miniblog/v1"
)

// UserController 是 user 模块在 Controller 层的实现，用来处理用户模块的请求.
type UserController struct {
	b biz.IBiz
}

// New 创建一个 user controller.
func New(ds store.IStore) *UserController {
	return &UserController{b: biz.NewBiz(ds)}
}

func (ctrl *UserController) Create(c *gin.Context) {
	log.C(c).Infow("Create user function called")

	var r v1.CreateUserRequest
	if err := c.ShouldBind(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)

		return
	}

	if err := ctrl.b.Users().Create(c, &r); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

}
