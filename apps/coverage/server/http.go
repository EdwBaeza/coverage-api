package server

import (
	"github.com/edwbaeza/coverage-api/apps/coverage/server/handlers/health"
	purchaseHandlers "github.com/edwbaeza/coverage-api/apps/coverage/server/handlers/purchase"
	statusHandlder "github.com/edwbaeza/coverage-api/apps/coverage/server/handlers/purchase/status"
	userHandlers "github.com/edwbaeza/coverage-api/apps/coverage/server/handlers/user"
	"github.com/edwbaeza/coverage-api/apps/coverage/server/middlewares"
	purchaseInfraestructure "github.com/edwbaeza/coverage-api/src/purchase/infraestructure"
	userInfraestructure "github.com/edwbaeza/coverage-api/src/user/infraestructure"
	"github.com/gin-gonic/gin"
)

// RegisterRouter wih gin context
func RegisterRouter(engine *gin.Engine) {
	purchaseRespository := purchaseInfraestructure.NewMongoRepository()
	userRepository := userInfraestructure.NewMongoRepository()

	engine.Use(middlewares.ErrorMiddleware())

	apiGroup := engine.Group("/api")
	authorized := apiGroup.Group("/")
	authorized.Use(middlewares.JwtAuthMiddleware(userRepository))

	authorized.GET("/purchases", purchaseHandlers.ListPurchasesHandler(purchaseRespository))
	authorized.GET("/purchases/:id", purchaseHandlers.FindPurchaseHandler(purchaseRespository))
	authorized.POST("/purchases", purchaseHandlers.CreatePurchaseHandler(purchaseRespository))
	authorized.PUT("/purchases/:id/status", statusHandlder.UpdatePurchaseStatusHandler(purchaseRespository))

	apiGroup.GET("/health_check", health.CheckHandler())
	apiGroup.POST("/users/tokens", userHandlers.TokenHandler(userRepository))
	apiGroup.POST("/users", userHandlers.RegistrationHandler(userRepository))
}
