package main

import (
	"errors"
	"github.com/go-redis/redis"
	"time"
)

type RedisBackend struct {
	client *redis.Client
}

func newRedisBackend(client *redis.Client) RedisBackend {
	return RedisBackend{client: client}
}
func (slf RedisBackend) Len(uuid UUID) (CNT, error) {
	result := slf.client.LLen(GetKeyFromUUID(uuid))
	return CNT(result.Val()), result.Err()
}
func (slf RedisBackend) Get(uuid UUID, cnt CNT) ([]MSG, error) {
	if cnt == 0 {
		return nil, errors.New("Invalid Argument zero. ")
	}
	if cnt < 0 {
		cnt = 0
	}
	result := slf.client.LRange(GetKeyFromUUID(uuid), 0, int64(cnt-1))
	msgs := make([]MSG, len(result.Val()))
	for i := 0; i < len(result.Val()); i++ {
		msg, err := GetMSGFromKey(result.Val()[i])
		if err != nil {
			println(err.Error())
			continue
		}
		msgs[i] = *msg
	}
	return msgs, result.Err()
}
func (slf RedisBackend) Write(dst UUID, msg *MSG) error {
	msg.TimeStamp = TimeStamp(time.Now().UnixNano())
	return slf.client.RPush(GetKeyFromUUID(dst), GetKeyFromMSG(msg)).Err()
}
func (slf RedisBackend) Skip(uuid UUID, cnt CNT) (CNT, error) {
	if cnt == 0 {
		return -1, errors.New("Invalid Argument zero. ")
	}
	t1, err := slf.Len(uuid)
	if err != nil {
		return -1, err
	}
	if cnt < 0 {
		cnt = t1
	}
	result := slf.client.LTrim(GetKeyFromUUID(uuid), int64(cnt), -1)
	if result.Err() != nil {
		return -1, result.Err()
	}
	t2, err := slf.Len(uuid)
	if err != nil {
		return -1, err
	}
	return t1 - t2, nil
}
