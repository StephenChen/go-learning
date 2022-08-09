package gock_demo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// ReqParam API 请求参数
type ReqParam struct {
	X int `json:"x"`
}

// Result API 返回结果
type Result struct {
	Value int `json:"value"`
}

func GetResultByAPI(x, y int) int {
	p := &ReqParam{X: x}
	b, _ := json.Marshal(p)

	// 调用其他服务的 API
	resp, err := http.Post(
		"http://your-api.com/post",
		"application/json",
		bytes.NewBuffer(b),
	)
	if err != nil {
		return -1
	}
	body, _ := io.ReadAll(resp.Body)
	var ret Result
	if err := json.Unmarshal(body, &ret); err != nil {
		return -1
	}

	return ret.Value + y
}
