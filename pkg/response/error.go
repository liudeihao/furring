package response

import (
    "fmt"
    "log"
    "net/http"
    "runtime"
    "strings"

    "github.com/gin-gonic/gin"
)

func trace(msg string) string {
    var pcs [32]uintptr
    n := runtime.Callers(0, pcs[:]) // skip first 3 caller

    var str strings.Builder
    str.WriteString(msg + "\nTraceback:")
    for _, pc := range pcs[:n] {
        fn := runtime.FuncForPC(pc)
        file, line := fn.FileLine(pc)
        str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
    }
    return str.String()
}

func errorResponse(c *gin.Context, code int, msg string) {
    c.AbortWithStatusJSON(code, gin.H{
        "error": msg,
    })
}

func Internal(c *gin.Context, err error) {
    log.Printf("服务器内部错误：%+v", err)
    log.Printf("%s\n\n", trace(err.Error()))
    errorResponse(c, http.StatusInternalServerError, "服务器内部错误")
}

func BadRequest(c *gin.Context, msg any) {
    switch msg.(type) {
    case error:
        errorResponse(c, http.StatusBadRequest, "参数错误："+msg.(error).Error())
    case string:
        errorResponse(c, http.StatusBadRequest, msg.(string))
    default:
        errorResponse(c, http.StatusBadRequest, "参数错误")
    }
}

func Unauthorized(c *gin.Context, msg string) {
    errorResponse(c, http.StatusUnauthorized, msg)
}

func NotFound(c *gin.Context, msg string) {
    errorResponse(c, http.StatusNotFound, msg+": 资源不存在")
}

func OK(c *gin.Context, data gin.H) {
    c.JSON(http.StatusOK, data)
}
