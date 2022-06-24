package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Bundary(c *gin.Context) {
    /* Generate Access Log */
    
    /* Request */
    request_id := uuid.New().String()
    c.Set("request_id", request_id)
    c.Header("x-request-id", request_id)
    c.Next()
    /* Response */
}
