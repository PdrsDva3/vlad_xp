package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"vlad_xp/docs"
	"vlad_xp/internal/delivery/handlers"
	"vlad_xp/internal/delivery/middleware"
	"vlad_xp/internal/repository/user"
	"vlad_xp/internal/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start(db *sqlx.DB) {
	r := gin.Default()
	r.ForwardedByClientIP = true
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRepo := user.InitUserRepo(db)
	userService := service.InitUserService(userRepo)
	userHandler := handlers.InitUserHandler(userService)

	userRouter := r.Group("/user")

	userRouter.POST("/create", userHandler.Create)
	userRouter.GET("/:id", userHandler.GetUser)
	userRouter.DELETE("/delete/:id", userHandler.Delete)
	userRouter.POST("/login", userHandler.Login)

	r.Use(middleware.CORSMiddleware())

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("error running client: %v", err.Error()))
	}
}
