package handlers

import (
	"net/http"
	"time"

	"github.com/Fameeeee/go-cloud-native/internal/database" // **แก้ path ตามชื่อ mod ของคุณ**
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RegisterRoutes(r *gin.Engine, rdb *redis.Client) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Endpoint สำหรับเก็บข้อมูลลง Redis
	r.GET("/set/:key/:value", func(c *gin.Context) {
		key := c.Param("key")
		val := c.Param("value")

		err := rdb.Set(database.Ctx, key, val, 10*time.Minute).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "saved", "key": key, "value": val})
	})

	// Endpoint สำหรับดึงข้อมูลจาก Redis
	r.GET("/get/:key", func(c *gin.Context) {
		key := c.Param("key")
		val, err := rdb.Get(database.Ctx, key).Result()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "key not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"key": key, "value": val})
	})
}
