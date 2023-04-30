/*
 * @Author: lwnmengjing
 * @Date: 2021/6/10 3:39 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/10 3:39 下午
 */

package storage

import (
	"github.com/go-admin-team/go-admin-core/storage"
	"log"

	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/config"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/captcha"
)

var CacheAdapter storage.AdapterCache
var Err error

// Setup 配置storage组件
func Setup() {
	//4. 设置缓存

	CacheAdapter, Err = config.CacheConfig.Setup()
	RedisErr := CacheAdapter.Set("Ping", "PongPong", 3600)
	if RedisErr != nil {
		log.Fatalf("cache setup error, %s\n", Err.Error())
		return
	}
	if Err != nil {
		log.Fatalf("cache setup error, %s\n", Err.Error())
	}
	sdk.Runtime.SetCacheAdapter(CacheAdapter)
	//5. 设置验证码store
	captcha.SetStore(captcha.NewCacheStore(CacheAdapter, 600))

	//6. 设置队列
	if !config.QueueConfig.Empty() {
		if q := sdk.Runtime.GetQueueAdapter(); q != nil {
			q.Shutdown()
		}
		queueAdapter, err := config.QueueConfig.Setup()
		if err != nil {
			log.Fatalf("queue setup error, %s\n", err.Error())
		}
		sdk.Runtime.SetQueueAdapter(queueAdapter)
		defer func() {
			go queueAdapter.Run()
		}()
	}

	//7. 设置分布式锁
	if !config.LockerConfig.Empty() {
		lockerAdapter, err := config.LockerConfig.Setup()
		if err != nil {
			log.Fatalf("locker setup error, %s\n", err.Error())
		}
		sdk.Runtime.SetLockerAdapter(lockerAdapter)
	}
}
