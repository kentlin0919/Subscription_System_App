package http

import "github.com/gin-gonic/gin"

// NewRouter wires HTTP routes and handlers using Gin.
func NewRouter(userHandler *UserHandler) *gin.Engine {
	router := gin.New()

	// 使用 Recovery middleware 確保 panic 不會終止 HTTP 服務。
	router.Use(gin.Recovery())

	if userHandler != nil {
		// 掛載使用者建立端點，後續可擴充其他 REST 資源。
		router.POST("/users", userHandler.HandleCreate)
	}

	return router
}
