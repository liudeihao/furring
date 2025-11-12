package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInternalServerErrorResponse(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/api", nil)

	InternalServerErrorResponse(c, assert.AnError)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "服务器内部错误", resp["error"])
}

func TestBadRequestResponse_WithoutDetails(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/api", nil)
	BadRequestResponse(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var resp map[string]string
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "请求参数有误", resp["error"])
}

func TestBadRequestResponse_WithDetailsString(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/api", nil)
	BadRequestResponse(c, "用户不存在")
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var resp map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "请求参数有误", resp["error"])
	assert.Equal(t, "用户不存在", resp["details"])
}
