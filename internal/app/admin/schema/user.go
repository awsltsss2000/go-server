package schema

import (
	"encoding/json"
	"go-server/internal/app/schema"
	"go-server/pkg/time"
)

// User 用户对象
type User struct {
	ID        uint64        `json:"id,string"`                             // 唯一标识
	UserName  string        `json:"user_name" binding:"required"`          // 用户名
	RealName  string        `json:"real_name" binding:"required"`          // 真实姓名
	Password  string        `json:"password"`                              // 密码
	Phone     string        `json:"phone"`                                 // 手机号
	Email     string        `json:"email"`                                 // 邮箱
	Status    int           `json:"status" binding:"required,max=2,min=1"` // 用户状态(1:启用 2:停用)
	Creator   uint64        `json:"creator"`                               // 创建者
	CreatedAt time.JSONTime `json:"created_at"`                            // 创建时间
}

func (u User) MarshalJSON() ([]byte, error) {
	if u.ID == 0 {
		return []byte("{}"), nil
	}
	type JSONUser User
	return json.Marshal(JSONUser(u))
}

// UserQueryParam 查询条件
type UserQueryParam struct {
	schema.PaginationParam
	UserName   string `form:"user_name"`   // 用户名
	QueryValue string `form:"query_value"` // 模糊查询
	Status     int    `form:"status"`      // 用户状态(1:启用 2:停用)
}

// UserQueryOptions 查询可选参数项
type UserQueryOptions struct {
	OrderFields  []*schema.OrderField
	SelectFields []string
}

type UserQueryResult struct {
	Data       Users
	PageResult *schema.PaginationResult
}

// Users 用户对象列表
type Users []*User
