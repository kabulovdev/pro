package main

import (
	"gitlab.com/pro/exam_api/api"
	"gitlab.com/pro/exam_api/config"
	"gitlab.com/pro/exam_api/pkg/logger"
	"gitlab.com/pro/exam_api/services"
	r "gitlab.com/pro/exam_api/storage/redis"
	"github.com/casbin/casbin/v2/util"
	"fmt"
	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	"github.com/casbin/casbin/v2"
	"github.com/gomodule/redigo/redis"
)


func main() {
	var (
		casbinEnforcer *casbin.Enforcer
	)
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "exam_api")

	//psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	//	cfg.PostgresHost,
	//	cfg.PostgresPort,
	//	cfg.PostgresUser,
	//	cfg.PostgresPassword,
	//	cfg.PostgresDB,
	//)
	//enf, err := gormadapter.NewAdapter("postgres", psqlString, true)
	//if err != nil {
	//	log.Error("gorm adapter error", logger.Error(err))
	//	return
	//}
	fmt.Println("hi")
	casbinEnforcer, err := casbin.NewEnforcer(cfg.AuthConfigPath, "./config/casbin.csv")	
	if err != nil {
		log.Error("casbin enforcer error", logger.Error(err))
		return
	}
	fmt.Println("hi")
	err = casbinEnforcer.LoadPolicy()
	if err != nil {
		log.Error("casbin error load policy", logger.Error(err))
		return
	}
	fmt.Println("hi")
	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}
	fmt.Println("hi")
	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("keyMatch", util.KeyMatch)
	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("keyMatch3", util.KeyMatch3)
	fmt.Println("hi")
	pool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", cfg.RedisHost+":"+cfg.RedisPort)
		},
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
		Redis:          r.NewRedisRepo(pool),
		CasbinEnforcer: casbinEnforcer,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}