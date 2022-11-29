package api

import (
	v1 "gitlab.com/pro/exam_api/api/handlers/v1"
	"gitlab.com/pro/exam_api/config"
	"gitlab.com/pro/exam_api/pkg/logger"
	"gitlab.com/pro/exam_api/services"
	"github.com/gin-contrib/cors"
	jwthandler "gitlab.com/pro/exam_api/api/tokens"
	middleware "gitlab.com/pro/exam_api/api/middleware"
	"gitlab.com/pro/exam_api/storage/repo"
	"github.com/casbin/casbin/v2"

	_ "gitlab.com/pro/exam_api/api/docs" //swag

	"github.com/gin-gonic/gin"
	swaggerFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Redis     repo.InMemoryStorageI
	CasbinEnforcer  *casbin.Enforcer
}

// New ...
// @Description Created by Abduazim Kabulov
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host      54.238.26.123:9079
// @BasePath  /v1

func New(option Option) *gin.Engine {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowCredentials: true,
		AllowOrigins:     []string{},
	}))


	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	jwt := jwthandler.JWTHandler{
		SigninKey: option.Conf.SignKey,
		Log:       option.Logger,
	}

	
	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
		Redis:          option.Redis,
		JWTHandler:      jwt,
	})

	router.Use(middleware.NewAuth(option.CasbinEnforcer, jwt, config.Load()))


	api := router.Group("/v1")
	api.POST("/custumer/create", handlerV1.CreateCustumer)
	api.GET("/custumer/get/:id", handlerV1.GetCustumer)
	api.DELETE("/custumer/delete/:id", handlerV1.DeleteCustumer)
	api.PUT("/custumer/update", handlerV1.UpdateCustumer)
	api.PUT("/post/update",handlerV1.UpdatePost)
	api.POST("/post/create", handlerV1.CreatePost)
	api.POST("/reating/create",handlerV1.CreateReating)
	api.PUT("/reating/update",handlerV1.UpdateReating)
	api.DELETE("/reating/delete/:id",handlerV1.DeleteReating)
	api.DELETE("/post/delet/:id",handlerV1.DeletePost)
	api.GET("/post/allInfo/:id",handlerV1.GetPostAllInfo)
	api.GET("/reating/get/:id",handlerV1.GetReating)
	api.GET("/post/get/reatings/:id",handlerV1.GetPostReating)
	api.GET("/custumer/getList",handlerV1.GetListCustumers)
	api.GET("/post/get/:id",handlerV1.GetPost)
	api.GET("/post/get/reatings/avarage/:id",handlerV1.GetPostReatingNew)
	api.POST("/register",handlerV1.RegisterUser)
	api.PATCH("/verify/:email/:code",handlerV1.Verify)
	api.GET("/token", handlerV1.GetAccesToken)
	api.GET("/admin/login/:name/:password", handlerV1.LoginAdmin)
	api.GET("/moder/login/:name/:password", handlerV1.LoginModerator)
	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFile.Handler, url))
	return router

}
