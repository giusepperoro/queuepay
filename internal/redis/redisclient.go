package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

// Очередь клиентов
const clientEventsQueuePrefixKey = "client_events"

// Очередь из транзакций клиента
const clientQueuePrefixKey = "client_id"

type Queue interface {
	Get() (bool, int64, int64)
	Set(clientId int64, amount int64) error
}

type RedisQueue struct {
	client *redis.Client
}

func RedisNew() *RedisQueue {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	if err != nil {
		log.Println(err)
		log.Fatal("redis connect error")
	}
	return &RedisQueue{client: client}
}

func (r *RedisQueue) Get() (bool, int64, int64) {

	res := r.client.LPop(clientEventsQueuePrefixKey)
	if res.Err() != nil {
		return false, 0, 0
	}
	clientId, err := res.Int64()
	if err != nil {
		return false, 0, 0
	}

	clientQueueKey := fmt.Sprintf("%s:%d", clientQueuePrefixKey, clientId)
	amount, err := r.client.LPop(clientQueueKey).Int64()
	if err != nil {
		return false, 0, 0
	}
	log.Println("Got client")
	return true, amount, clientId
}

func (r *RedisQueue) Set(clientId, amount int64) error {
	clientQueueKey := fmt.Sprintf("%s:%d", clientQueuePrefixKey, clientId)
	log.Printf("in Set clientId: %d, amount %d", clientId, amount)
	res := r.client.LPush(clientEventsQueuePrefixKey, clientId)
	if res.Err() != nil {
		return fmt.Errorf("unable to save to queue %s:%w", clientQueueKey, res.Err())
	}

	res = r.client.LPush(clientQueueKey, amount)
	if res.Err() != nil {
		return fmt.Errorf("unable to save client %d to set of availible clients %s:%w", clientId, clientEventsQueuePrefixKey, res.Err())
	}
	return nil
}
