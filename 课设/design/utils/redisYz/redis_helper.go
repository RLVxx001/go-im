package redisYz

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"time"
)

var key = "verification_code_123"
var ctx = context.Background()
var client *redis.Client

// 初始化Redis客户端（同原代码）
func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "192.168.11.15:6380",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})
}

// 生成6位随机验证码
func generateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, 6)
	for i := range code {
		code[i] = byte(rand.Intn(10) + 48) // 生成0-9之间的数字
	}
	return string(code)
}

// 设置验证码，并设置过期时间为1分钟
func SetVerificationCode() error {
	code := generateVerificationCode()
	expiration := time.Minute
	err := client.Set(ctx, key, code, expiration).Err()
	if err != nil {
		return err
	}
	fmt.Printf("Set verification code for %s: %s (expires after %s)\n", key, code, expiration)
	return nil
}

// 获取验证码，如果验证码不存在或过期，则重新设置并返回新验证码
func GetVerificationCode() (string, error) {
	val, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		// 键不存在，重新设置验证码
		if err := SetVerificationCode(); err != nil {
			return "", err
		}
		// 重新获取（实际上可以直接返回上面生成的验证码，但为了保持一致性，这里还是再获取一次）
		val, err = client.Get(ctx, key).Result()
		if err != nil {
			return "", err
		}
	} else if err != nil {
		// 其他错误
		return "", err
	}
	// 检查验证码是否过期（这里Redis已经自动处理了过期，所以实际上不需要再检查）
	// 但如果需要更精细的控制，可以获取TTL并判断
	return val, nil
}
