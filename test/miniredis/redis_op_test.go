package miniredis_demo

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

func TestDoSomethingWithRedis(t *testing.T) {
	// mock 一个 redis server
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	// 准备数据
	s.Set("cxy", "github.com/stephenchen")
	s.SAdd(KeyValidWebsite, "cxy")

	// 连接 mock 的 redis server
	rdb := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	// 调用函数
	ok := DoSomethingWithRedis(rdb, "cxy")
	if !ok {
		t.Fatal()
	}

	// 可以手动检查 redis 中的值是否符合预期
	if got, err := s.Get("blog"); err != nil || got != "https://github.com/stephenchen" {
		t.Fatalf("'blog' has the wrong value")
	}

	// 也可以使用帮助工具检查
	s.CheckGet(t, "blog", "https://github.com/stephenchen")

	// 过期检查
	s.FastForward(5 * time.Second)
	if s.Exists("blog") {
		t.Fatalf("'blog' should not have existed anymore")
	}
}
