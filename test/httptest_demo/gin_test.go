package httptest_demo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_helloHandler(t *testing.T) {
	tests := []struct {
		name   string
		param  string
		expect string
	}{
		{"base case", `{"name": "cxy"}`, "hello cxy"},
		{"bad case", "", "we need a name"},
	}

	r := SetupRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock 一个 HTTP 请求
			req := httptest.NewRequest(
				"POST",
				"/hello",
				strings.NewReader(tt.param),
			)

			// mock 一个响应记录器
			w := httptest.NewRecorder()

			// 让 server 端处理 mock 请求并记录返回的响应内容
			r.ServeHTTP(w, req)

			// 校验状态码是否需符合预期
			assert.Equal(t, http.StatusOK, w.Code)

			// 解析并检验响应内容是否符合预期
			var resp map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			assert.Nil(t, err)
			assert.Equal(t, tt.expect, resp["msg"])
		})
	}
}
