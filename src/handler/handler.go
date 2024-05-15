package handler

import "github.com/gin-gonic/gin"

type AuthHandlerInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}