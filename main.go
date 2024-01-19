package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"bee_http_game/service"
)

func main() {
	r := gin.Default()

	r2 := gin.New()
	r2.Use(gin.Logger())
	r2.Use(gin.Recovery())
	//gin.SetMode(setting.RunMode)
	apiv1 := r2.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", service.GetEnv)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
	return r

	r.GET("/tags", service.GetEnv)

	//r.GET("/env/:id", func(c *gin.Context) {
	//	service.GetEnv
	//	//id := c.GetInt("id")
	//	//c.String(http.StatusOK, "Hello %s", id)
	//})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

