package server

import (
	"github.com/gin-gonic/gin"
	"go-docker-test/cache"
	"go-docker-test/model"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// 路由
	v1 := r.Group("")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1.GET("list", queryUser)

		v1.POST("set",redisSet)
		v1.GET("get",redisGet)
	}
	return r
}

func queryUser(c *gin.Context) {
	var users []model.User
	err := model.DB.
		Model(&model.User{}).
		Find(&users).
		Limit(10).
		Error
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, users)
	}
}

func redisSet(c *gin.Context) {
	key, _ := c.GetQuery("name")

	cache.RedisClient.Set("name", key, 0)

	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func redisGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": cache.RedisClient.Get("name").Val(),
	})

}
