package http

import "github.com/gin-gonic/gin"

// SetupRouter はルーティングを設定してGinエンジンを返す
func SetupRouter(tweetHandler *TweetHandler) *gin.Engine {
	router := gin.Default()

	// API v1
	api := router.Group("/api")
	{
		api.POST("/tweets", tweetHandler.CreateTweet)
	}

	return router
}
