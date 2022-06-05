package app

import (
	"fmt"
	"go-server/internal/app/admin/api"
	"go-server/internal/app/admin/dao/user"
	"go-server/internal/app/admin/service"
	dbUtil "go-server/internal/app/db"
	"go-server/internal/app/router"
	"net/http"
	"time"
)

func Init() {
	InitConfig()
	cfg := GetConfig()

	InitRedis()

	db := InitDB()

	transRepo := &dbUtil.Trans{DB: db}
	userRepo := &user.Repo{DB: db}
	userSrv := &service.UserSrv{
		TransRepo: transRepo,
		UserRepo:  userRepo,
	}
	userAPI := &api.UserApi{UserSrv: userSrv}
	r := &router.Router{UserApi: userAPI}

	engine := InitGin(r)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	server := &http.Server{
		Addr:              addr,
		Handler:           engine,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	// go func() {
	// 	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 	}
	// }()
	//
	// shutdown.NewHook().Close(
	// 	// 关闭 http server
	// 	func() {
	// 		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	// 		defer cancel()
	//
	// 		if err := server.Shutdown(ctx); err != nil {
	// 		}
	// 	},
	//
	// 	// 关闭 db
	// 	func() {
	//
	// 		if db != nil {
	// 			sqlDB, err := db.DB()
	// 			if err != nil {
	//
	// 			}
	// 			err = sqlDB.Close()
	// 			if err != nil {
	//
	// 			}
	// 		}
	// 	},
	// )

}
