package user

import (
	"context"
	"go-server/internal/app/admin/schema"
	"go-server/internal/app/db"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func GetUserDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return db.GetDBWithModel(ctx, defDB, new(User))
}

type SchemaUser schema.User

func (u SchemaUser) ToUser() *User {
	item := new(User)
	copier.Copy(item, u)
	return item
}

type User struct {
	db.Model
	UserName string `gorm:"size:64;uniqueIndex;default:'';not null;"` // 用户名
	RealName string `gorm:"size:64;index;default:'';"`                // 真实姓名
	Password string `gorm:"size:40;default:'';"`                      // 密码
	Email    string `gorm:"size:255;"`                                // 邮箱
	Phone    string `gorm:"size:20;"`                                 // 手机号
	Status   int    `gorm:"index;default:0;"`                         // 状态(1:启用 2:停用)
	Creator  uint64 `gorm:""`                                         // 创建者
}

func (u *User) TableName() string {
	return "admin_user"
}

func (u *User) ToSchemaUser() *schema.User {
	item := new(schema.User)
	copier.Copy(item, u)
	return item
}

type Users []*User

func (u Users) ToSchemaUsers() []*schema.User {
	list := make([]*schema.User, len(u))
	for i, item := range u {
		list[i] = item.ToSchemaUser()
	}
	return list
}
