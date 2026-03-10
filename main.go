package main

import (
	"log"

	"github.com/Fameeeee/go-cloud-native/internal/database"
	"github.com/Fameeeee/go-cloud-native/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// เชื่อมต่อ Redis
	rdb, err := database.NewRedisClient()
	if err != nil {
		log.Fatalf("Fatal: %v", err)
	}
	log.Println("✅ Connected to Redis")

	// ตั้งค่า Gin
	r := gin.Default()

	// ลงทะเบียน API Routes
	handlers.RegisterRoutes(r, rdb)

	// เริ่มรัน Server
	log.Println("🚀 Server is running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
