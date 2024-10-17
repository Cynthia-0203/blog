package middlewares

import (
	"github.com/gin-gonic/gin"
)

type HttpHandlerFunc func(*gin.Context)