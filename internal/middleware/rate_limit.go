package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	lastRequest = make(map[string]time.Time)
	mu          sync.Mutex
)

func OneRequestPer10Seconds() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		fmt.Println(
			c.ClientIP(),
			c.Request.RemoteAddr,
			c.GetHeader("X-Forwarded-For"),
		)

		mu.Lock()
		last, exists := lastRequest[ip]

		if exists && time.Since(last) < 10*time.Second {
			wait := 10 - int(time.Since(last).Seconds())
			mu.Unlock()

			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":       "too many requests",
				"retry_after": wait,
			})
			c.Abort()
			return
		}

		lastRequest[ip] = time.Now()
		mu.Unlock()

		c.Next()
	}
}
