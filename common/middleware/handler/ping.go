package handler

import (
	"github.com/gin-gonic/gin"
	"go-admin/common/myKafka"
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

func MessageQue(c *gin.Context) {
	err := myKafka.SendMessage("hello kafka", "my-topic", storage.KafkaConn)
	if err != nil {
		log.Fatalf("kafka setup error, %s\n", storage.Err.Error())
		return
	}
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
