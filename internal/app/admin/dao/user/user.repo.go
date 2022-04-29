package user

import (
	"context"
	"go-server/internal/app/admin/schema"
	dbUtil "go-server/internal/app/db"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func (r *Repo) getQueryOption(opts ...schema.UserQueryOptions) schema.UserQueryOptions {
	var opt schema.UserQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (r *Repo) List(ctx context.Context, params *schema.UserQueryParam, opts ...schema.UserQueryOptions) (*schema.UserQueryResult, error) {
	opt := r.getQueryOption(opts...)

	db := GetUserDB(ctx, r.DB)
	if v := params.UserName; v != "" {
		db = db.Where("user_name=?", v)
	}
	if v := params.Status; v > 0 {
		db = db.Where("status=?", v)
	}
	if v := params.QueryValue; v != "" {
		v = "%" + v + "%"
		db = db.Where("user_name like ? or real_name like ?", v, v)
	}
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(dbUtil.ParseOrder(opt.OrderFields))
	}

	var list Users
	pr, err := dbUtil.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.UserQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaUsers(),
	}
	return qr, nil
}

func (r *Repo) Retrieve(ctx context.Context, id uint64, opts ...schema.UserQueryParam) (*schema.User, error) {
	item := new(User)
	ok, err := dbUtil.FindOne(ctx, GetUserDB(ctx, r.DB).Where("id=?", id), item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item.ToSchemaUser(), nil
}

func (r *Repo) Create(ctx context.Context, item *schema.User) (uint64, error) {
	user := SchemaUser(*item).ToUser()
	result := GetUserDB(ctx, r.DB).Create(user)
	return user.ID, errors.WithStack(result.Error)
}

func (r *Repo) Update(ctx context.Context, id uint64, item *schema.User) error {
	user := SchemaUser(*item).ToUser()
	result := GetUserDB(ctx, r.DB).Where("id=?", id).Updates(user)
	return errors.WithStack(result.Error)
}

func (r *Repo) Delete(ctx context.Context, id uint64) error {
	result := GetUserDB(ctx, r.DB).Where("id=?", id).Delete(User{})
	return errors.WithStack(result.Error)
}
