package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	pb "github.com/betaiota/grpchat/proto/chatpb"
	"github.com/redis/go-redis/v9"
)

const (
	chatMessagesKey = "chat:messages"
	userMessagesKey = "chat:user:%s:messages"
	maxMessages     = 1000
	messageTTL      = 7 * 24 * time.Hour
)

type RedisStorage struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStorage(addr string) (*RedisStorage, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	log.Printf("Connected to Redis at %s", addr)

	return &RedisStorage{
		client: rdb,
		ctx:    ctx,
	}, nil
}

func (r *RedisStorage) Close() error {
	return r.client.Close()
}

func (r *RedisStorage) StoreMessage(msg *pb.ChatMessage) error {
	msgJSON, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	key := chatMessagesKey
	if err := r.client.ZAdd(r.ctx, key, redis.Z{
		Score:  float64(msg.Timestamp),
		Member: string(msgJSON),
	}).Err(); err != nil {
		return fmt.Errorf("failed to store message in global set: %w", err)
	}

	userKey := fmt.Sprintf(userMessagesKey, msg.Username)
	if err := r.client.ZAdd(r.ctx, userKey, redis.Z{
		Score:  float64(msg.Timestamp),
		Member: string(msgJSON),
	}).Err(); err != nil {
		return fmt.Errorf("failed to store message in user set: %w", err)
	}

	r.client.Expire(r.ctx, userKey, messageTTL)

	if err := r.client.ZRemRangeByRank(r.ctx, key, 0, -maxMessages-1).Err(); err != nil {
		log.Printf("Warning: failed to trim messages: %v", err)
	}

	if err := r.client.ZRemRangeByRank(r.ctx, userKey, 0, -maxMessages-1).Err(); err != nil {
		log.Printf("Warning: failed to trim user messages: %v", err)
	}

	return nil
}

func (r *RedisStorage) GetRecentMessages(limit int64) ([]*pb.ChatMessage, error) {
	messages, err := r.client.ZRevRange(r.ctx, chatMessagesKey, 0, limit-1).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve messages: %w", err)
	}

	result := make([]*pb.ChatMessage, 0, len(messages))
	for _, msgJSON := range messages {
		var msg pb.ChatMessage
		if err := json.Unmarshal([]byte(msgJSON), &msg); err != nil {
			log.Printf("Warning: failed to unmarshal message: %v", err)
			continue
		}
		result = append(result, &msg)
	}

	return result, nil
}

func (r *RedisStorage) GetUserMessages(username string, limit int64) ([]*pb.ChatMessage, error) {
	userKey := fmt.Sprintf(userMessagesKey, username)

	messages, err := r.client.ZRevRange(r.ctx, userKey, 0, limit-1).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user messages: %w", err)
	}

	result := make([]*pb.ChatMessage, 0, len(messages))
	for _, msgJSON := range messages {
		var msg pb.ChatMessage
		if err := json.Unmarshal([]byte(msgJSON), &msg); err != nil {
			log.Printf("Warning: failed to unmarshal message: %v", err)
			continue
		}
		result = append(result, &msg)
	}

	return result, nil
}

func (r *RedisStorage) GetMessageCount() (int64, error) {
	count, err := r.client.ZCard(r.ctx, chatMessagesKey).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to get message count: %w", err)
	}
	return count, nil
}
