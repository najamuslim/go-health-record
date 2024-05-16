package handler

import (
	"github.com/gin-gonic/gin"
)

type AuthHandlerInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type NurseHandlerInterface interface {
  RegisterNurse(c *gin.Context)
	// UpdateNurse(c *gin.Context)
	// DeleteNurse(c *gin.Context)
	// GetNurses(c *gin.Context)
}