package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID   uint64
	Name string
}

func main() {
	users := []User{
		{ID: 123, Name: "zs"},
		{ID: 456, Name: "ls"},
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, c.DefaultQuery("wechat", "cxy"))
		c.JSON(200, c.QueryArray("media"))
	})
	r.GET("/map", func(c *gin.Context) {
		c.JSON(200, c.QueryMap("ids"))
	})
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})
	r.GET("/users/*id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "The user id is %s", id)
	})
	r.POST("/", func(c *gin.Context) {
		wechat := c.PostForm("wechat")
		c.String(200, wechat)
	})
	r.DELETE("/users/123", func(c *gin.Context) {
	})
	r.PUT("/users/123", func(c *gin.Context) {
	})
	r.PATCH("/users/123", func(c *gin.Context) {
	})
	//r.Any("/", func(context *gin.Context) {
	//
	//}, func(context *gin.Context) {
	//
	//})

	// V1 version API
	v1Group := r.Group("/v1")
	v1Group.GET("/users", func(c *gin.Context) {
		c.String(200, "/v1/users")
	})
	v1Group.GET("/products", func(c *gin.Context) {
		c.String(200, "/v1/products")
	})
	// V2 version API
	v2Group := r.Group("/v2")
	v2Group.GET("/users", func(c *gin.Context) {
		c.String(200, "/v2/users")
	})
	v2Group.GET("/products", func(c *gin.Context) {
		c.String(200, "/v2/products")
	})

	r.Run(":8080")

}
