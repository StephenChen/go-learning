package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"github.com/go-playground/validator/v10"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Rest GET, POST, PUT, PATCH, DELETE, OPTIONS
func Rest(router *gin.Engine) {
	if router == nil {
		router = gin.Default()
	}

	getting := func(ctx *gin.Context) { ctx.String(200, "restGet") }
	posting := func(ctx *gin.Context) { ctx.String(200, "restPost") }
	putting := func(ctx *gin.Context) { ctx.String(200, "restPut") }
	deleting := func(ctx *gin.Context) { ctx.String(200, "restDelete") }
	patching := func(ctx *gin.Context) { ctx.String(200, "restPatch") }
	head := func(ctx *gin.Context) { ctx.String(200, "restHead") }
	options := func(ctx *gin.Context) { ctx.String(200, "restOptions") }

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)
}

// RouteParam 路由参数
func RouteParam(router *gin.Engine) {
	if router == nil {
		router = gin.Default()
	}

	// 匹配 /user/chen 不匹配 /user/ 或 /user
	router.GET("/user/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "Hello %s", name)
	})

	// 匹配 /user/chen/ 及 /user/chen/send
	// 如果没有其他路由匹配 /user/chen, 将重定向至 /user/chen/
	router.GET("/user/:name/*action", func(ctx *gin.Context) {
		name := ctx.Param("name")
		action := ctx.Param("action")
		message := name + " is " + action
		ctx.String(http.StatusOK, message)
	})

	// 匹配的路由, 上下文保留路由定义
	router.POST("/user/:name/*action", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%v", ctx.FullPath() == "/user/:name/*action")
	})
}

func query(router *gin.Engine) {
	if router == nil {
		router = gin.Default()
	}

	router.GET("/welcome", func(ctx *gin.Context) {
		firstname := ctx.DefaultQuery("firstname", "Guest")
		lastname := ctx.Query("lastname")

		ctx.String(http.StatusOK, "hello %s %s", firstname, lastname)
	})
}

func multipart(router *gin.Engine) {
	if router == nil {
		router = gin.Default()
	}

	router.POST("/form_post", func(ctx *gin.Context) {
		message := ctx.PostForm("message")
		nick := ctx.DefaultPostForm("nick", "anonymous")

		ctx.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
}

func group(router *gin.Engine) {
	if router == nil {
		router = gin.Default()
	}

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/login", func(c *gin.Context) { c.String(200, "loginV1") })
		v1.GET("/submit", func(c *gin.Context) { c.String(200, "submitV1") })
		v1.GET("/read", func(c *gin.Context) { c.String(200, "readV1") })
	}
	// Simple group v2
	v2 := router.Group("/v2")
	{
		v2.GET("/login", func(c *gin.Context) { c.String(200, "loginV2") })
		v2.GET("/submit", func(c *gin.Context) { c.String(200, "submitV2") })
		v2.GET("/read", func(c *gin.Context) { c.String(200, "readV2") })
	}
}

func middleware() {
	// 创建一个默认的没有任何中间件的路由
	// gin.Default 默认已连接 Logger and Recovery 中间件
	r := gin.New()

	// 全局中间件

	// Logger 中间件将日志写到 gin.DefaultWriter 即使设置 GIN_MODE=release
	// 默认设置 gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery 中间件从任何 panic 恢复, 如果出现 panic, 会写一个 500 错误
	r.Use(gin.Recovery())
}

// 绑定为 JSON
func bind() {
	type Login struct {
		User     string `form:"user" json:"user" xml:"user" binding:"required"`
		Password string `form:"password" json:"password" xml:"password"`
	}

	r := gin.Default()

	// JSON 绑定实例 ({"user": "manu", "password": "123"})
	r.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
}

// 自定义验证器

// Booking 预定包含绑定和验证的数据
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func customValidator(r *gin.Engine) {
	if r == nil {
		r = gin.Default()
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("bookabledate", bookableDate)
	}

	r.GET("/bookable", getBookable)
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// XML, JSON, YAML, ProtoBuf 渲染
func render(r *gin.Engine) {
	if r == nil {
		r = gin.Default()
	}

	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Chen"
		msg.Message = "hey"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(http.StatusOK, data)
	})
}

// using BasicAuth
// 模拟些私有数据
var secrets = gin.H{
	"foo":    gin.H{"email": "qwe@qwe.com", "phone": "123123"},
	"austin": gin.H{"email": "asd@asd.com", "phone": "456456"},
	"lena":   gin.H{"email": "zxc@zxc.com", "phone": "789789"},
}

func basicAuth(r *gin.Engine) {
	if r == nil {
		r = gin.Default()
	}

	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"chen":   "chen",
	}))

	authorized.GET("/secrets", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})
}

// 多种服务

var g errgroup.Group

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "error": "welcome server 01"})
	})
	return e
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "error": "welcome server 02"})
	})
	return e
}

func multiServer() {
	server01 := http.Server{
		Addr:         ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server02 := http.Server{
		Addr:         ":8081",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		err := server01.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})
	g.Go(func() error {
		err := server02.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

// 正常启动停止
func startAndStop() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "welcome gin server")
	})

	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号超时5秒, 正常关闭服务器
	quit := make(chan os.Signal)
	// kill(无参) 默认发送 syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL 但不能被捕获, 不需要添加
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// 上下文用于通知服务器它有5秒的时间完成当前正在处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}

// cookie
func cookieGetSet(r *gin.Engine) {
	if r == nil {
		r = gin.Default()
	}

	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/",
				"localhost", false, true)
		}

		fmt.Printf("Cookie value: %s\n", cookie)
	})
}
