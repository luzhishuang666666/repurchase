package handler

import (
	"github.com/gin-gonic/gin"
	"go-admin/common/storage"
	"log"
)

func Ping(c *gin.Context) {
	RedisErr := storage.CacheAdapter.Set("TEST", "OK", 36)
	if RedisErr != nil {
		log.Fatalf("cache setup error, %s\n", storage.Err.Error())
		return
	}
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
