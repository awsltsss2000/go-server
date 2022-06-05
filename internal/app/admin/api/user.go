package api

import (
	"go-server/internal/app/admin/schema"
	"go-server/internal/app/admin/service"
	schemaBase "go-server/internal/app/schema"
	"go-server/internal/pkg/ginx"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	UserSrv *service.UserSrv
}

func (u *UserApi) List(c *gin.Context) {
	params := new(schema.UserQueryParam)
	if err := ginx.ParseQuery(c, params); err != nil {
		ginx.ResError(c, err)
		return
	}

	// params.OnlyCount = true
	params.Pagination = true
	opt := schema.UserQueryOptions{
		// SelectFields: []string{"id", "user_name", "real_name", "password", "phone", "email"},
		OrderFields: []*schemaBase.OrderField{schemaBase.NewOrderField("id", schemaBase.OrderByASC)},
	}
	result, err := u.UserSrv.List(c, params, opt)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResPage(c, result.Data, result.PageResult)
}

func (u *UserApi) Retrieve(c *gin.Context) {
	userId := ginx.ParseParamID(c, "userId")
	item, err := u.UserSrv.Retrieve(c, userId)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.Success(c, item)
}

func (u *UserApi) Create(c *gin.Context) {
	item := new(schema.User)
	if err := ginx.ParseJSON(c, item); err != nil {
		ginx.ResError(c, err)
		return
	}
	result, err := u.UserSrv.Create(c, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.Success(c, result)
}

func (u *UserApi) BulkCreate(c *gin.Context) {
	item := new(schema.Users)
	if err := ginx.ParseJSON(c, item); err != nil {
		ginx.ResError(c, err)
		return
	}
	result, err := u.UserSrv.BulkCreate(c, *item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.Success(c, result)
}

func (u *UserApi) Update(c *gin.Context) {
	userId := ginx.ParseParamID(c, "userId")
	item := new(schema.User)
	if err := ginx.ParseJSON(c, item); err != nil {
		ginx.ResError(c, err)
		return
	}
	result, err := u.UserSrv.Update(c, userId, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.Success(c, result)
}

func (u *UserApi) Delete(c *gin.Context) {
	userId := ginx.ParseParamID(c, "userId")
	err := u.UserSrv.Delete(c, userId)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.Success(c, nil)
}
