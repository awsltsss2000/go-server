package dao

import (
	"go-server/internal/app/admin/dao/user"
	"go-server/internal/app/db"
)

type (
	TransRepo = db.Trans
	UserRepo  = user.Repo
)
