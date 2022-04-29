package service

import (
	"context"
	"go-server/internal/app/admin/dao"
	"go-server/internal/app/admin/schema"
	schemaBase "go-server/internal/app/schema"
)

type UserSrv struct {
	TransRepo *dao.TransRepo
	UserRepo  *dao.UserRepo
}

func (u *UserSrv) List(ctx context.Context, params *schema.UserQueryParam, opts ...schema.UserQueryOptions) (*schema.UserQueryResult, error) {
	return u.UserRepo.List(ctx, params, opts...)
}

func (u *UserSrv) Retrieve(ctx context.Context, id uint64) (*schema.User, error) {
	item, err := u.UserRepo.Retrieve(ctx, id)
	if err != nil {
		return nil, err
	} else if item == nil {
		// return not found error
	}
	return item, nil
}

func (u *UserSrv) Create(ctx context.Context, item *schema.User) (*schemaBase.IDResult, error) {
	id, err := u.UserRepo.Create(ctx, item)
	return &schemaBase.IDResult{
		ID: id,
	}, err
}

// BulkCreate 主要是为了测试事务
func (u *UserSrv) BulkCreate(ctx context.Context, item []*schema.User) (*schemaBase.IDSResult, error) {
	var ids []uint64
	err := u.TransRepo.Exec(ctx, func(ctx context.Context) error {
		for _, uItem := range item {
			id, err := u.UserRepo.Create(ctx, uItem)
			if err != nil {
				return err
			}
			ids = append(ids, id)
		}
		return nil
	})

	return &schemaBase.IDSResult{
		IDs: ids,
	}, err
}

func (u *UserSrv) Update(ctx context.Context, id uint64, item *schema.User) (*schemaBase.IDResult, error) {
	err := u.UserRepo.Update(ctx, id, item)
	return &schemaBase.IDResult{
		ID: id,
	}, err
}

func (u *UserSrv) Delete(ctx context.Context, id uint64) error {
	return u.UserRepo.Delete(ctx, id)
}
