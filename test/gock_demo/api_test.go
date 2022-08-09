package gock_demo

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestGetResultByAPI(t *testing.T) {
	defer gock.Off() // 测试执行后刷新挂起的 mock

	// mock 请求外部 api 时传参 x=1 返回 100
	gock.New("http://your-api.com").
		Post("/post").
		MatchType("json").
		JSON(map[string]int{"x": 1}).
		Reply(200).
		JSON(map[string]int{"value": 100})

	// 调用业务函数
	res := GetResultByAPI(1, 1)
	// 校验返回结果是否符合预期
	assert.Equal(t, 101, res)

	gock.New("http://your-api.com").
		Post("/post").
		MatchType("json").
		JSON(map[string]int{"x": 2}).
		Reply(200).
		JSON(map[string]int{"value": 200})
	res = GetResultByAPI(2, 2)
	assert.Equal(t, 202, res)

	assert.True(t, gock.IsDone()) // 断言 mock 被触发
}

func TestGetResultByAPIWithTable(t *testing.T) {
	defer gock.Off() // 测试执行后刷新挂起的 mock

	// mock 请求外部 api 时传参 x=1 返回 100
	gock.New("http://your-api.com").
		Post("/post").
		MatchType("json").
		JSON(map[string]int{"x": 1}).
		Reply(200).
		JSON(map[string]int{"value": 100})
	// mock 请求外部 api 时传参 x=2 返回 200
	gock.New("http://your-api.com").
		Post("/post").
		MatchType("json").
		JSON(map[string]int{"x": 2}).
		Reply(200).
		JSON(map[string]int{"value": 200})

	tests := []struct {
		name   string
		params [2]int
		expect int
	}{
		{"case1", [2]int{1, 1}, 101},
		{"case2", [2]int{2, 2}, 202},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := GetResultByAPI(tt.params[0], tt.params[1])
			assert.Equal(t, tt.expect, res)
		})
	}

	assert.True(t, gock.IsDone()) // 断言 mock 被触发
}
