package rest

import (
	"os"

	"github.com/gin-contrib/cors"
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
	api "github.com/saipulmuiz/krplus/service"
	middlewares "github.com/saipulmuiz/krplus/service/middleware"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	userUsecase        api.UserUsecase
	creditUsecase      api.CreditUsecase
	transactionUsecase api.TransactionUsecase
}

func CreateHandler(
	userUsecase api.UserUsecase,
	creditUsecase api.CreditUsecase,
	transactionUsecase api.TransactionUsecase,
) *gin.Engine {
	obj := Handler{
		userUsecase:        userUsecase,
		creditUsecase:      creditUsecase,
		transactionUsecase: transactionUsecase,
	}

	var maxSize int64 = 1024 * 1024 * 10 //10 MB
	logger := log.New()
	r := gin.Default()
	mainRouter := r.Group("/v1")

	gin.SetMode(gin.DebugMode)
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	corsconfig := cors.DefaultConfig()
	corsconfig.AllowAllOrigins = true
	corsconfig.AddAllowHeaders("Authorization")
	r.Use(cors.New(corsconfig))
	r.Use(limits.RequestSizeLimiter(maxSize))
	r.Use(middlewares.ErrorHandler(logger))

	mainRouter.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	mainRouter.POST("/register", obj.Register)
	mainRouter.POST("/login", obj.login)

	authorizedRouter := mainRouter.Group("/")
	authorizedRouter.Use(middlewares.Auth())
	{
		authorizedRouter.GET("/credits", obj.GetCredits)
	}

	return r
}
