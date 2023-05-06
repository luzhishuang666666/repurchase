package middleware

import (
	"github.com/bwmarrin/snowflake"
	"log"
)

func GenerateDistributedId() (string, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

	// 生成一个唯一 ID
	id := node.Generate().String()
	log.Printf("generate Id :", id)
	return id, err
}
