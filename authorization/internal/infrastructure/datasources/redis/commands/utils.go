package commands

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

func GetStringValueIfExists(context context.Context, client *redis.Client, key string) (string, error) {
	cmd := client.Get(context, key)

	if cmd.Err() != nil {
		if !errors.Is(cmd.Err(), redis.Nil) {
			return "", cmd.Err()
		}
		return "", nil
	}

	return cmd.Result()
}

func GetValueIfExists[V any](context context.Context, client *redis.Client, key string) (*V, error) {
	str, err := GetStringValueIfExists(context, client, key)
	if err != nil {
		return nil, err
	}
	if str == "" {
		return nil, nil
	}

	var value V
	err = json.Unmarshal([]byte(str), &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func SetValue[V any](context context.Context, client *redis.Client, key string, value V, duration time.Duration) error {
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}
	client.Set(context, key, json, duration)
	return nil
}
