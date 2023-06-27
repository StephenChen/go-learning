package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("gin/api")
	//startAndStop()
	r := gin.Default()

	Rest(r)
	RouteParam(r)
	customValidator(r)
	render(r)
	basicAuth(r)

	_ = r.Run(":8080")
}
