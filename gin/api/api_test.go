package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRest(t *testing.T) {
	r := gin.Default()
	Rest(r)

	tests := []struct {
		name     string
		request  *http.Request
		status   int
		expected string
	}{
		{"GET", httptest.NewRequest("GET", "/someGet", nil), http.StatusOK, "restGet"},
		{"POST", httptest.NewRequest("POST", "/somePost", nil), http.StatusOK, "restPost"},
		{"PUT", httptest.NewRequest("PUT", "/somePut", nil), http.StatusOK, "restPut"},
		{"DELETE", httptest.NewRequest("DELETE", "/someDelete", nil), http.StatusOK, "restDelete"},
		{"PATCH", httptest.NewRequest("PATCH", "/somePatch", nil), http.StatusOK, "restPatch"},
		{"HEAD", httptest.NewRequest("HEAD", "/someHead", nil), http.StatusOK, "restHead"},
		{"OPTIONS", httptest.NewRequest("OPTIONS", "/someOptions", nil), http.StatusOK, "restOptions"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()

			// server 处理请求并返回响应内容
			r.ServeHTTP(recorder, tt.request)

			// 校验状态码
			assert.Equal(t, tt.status, recorder.Code)
			assert.Equal(t, tt.expected, recorder.Body.String())
		})
	}

}

func TestRouteParam(t *testing.T) {
	r := gin.Default()
	RouteParam(r)

	tests := []struct {
		name     string
		request  *http.Request
		status   int
		expected string
	}{
		{"/user/chen", httptest.NewRequest("GET", "/user/chen", nil), http.StatusOK, "Hello chen"},
		{"/user/chen/", httptest.NewRequest("GET", "/user/chen/", nil), http.StatusOK, "chen is /"},
		{"/user/chen/send", httptest.NewRequest("GET", "/user/chen/send", nil), http.StatusOK, "chen is /send"},
		{"/user/chen/send", httptest.NewRequest("POST", "/user/chen/send", nil), http.StatusOK, "true"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			r.ServeHTTP(recorder, tt.request)

			assert.Equal(t, tt.status, http.StatusOK)
			assert.Equal(t, tt.expected, recorder.Body.String())
		})
	}
}
