package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	h := NewHandler(db)
	h.RegisterHandlers(r)

	return r
}
