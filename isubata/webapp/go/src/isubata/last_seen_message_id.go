package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func lastSeenMessageIDKey(userID int64) string {
	return fmt.Sprintf("last_seen_message:%d", userID)
}

func setLastSeenMessageID(userID int64, channelID int64, messageID int64) error {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("HSET", lastSeenMessageIDKey(userID), channelID, messageID)
	return err
}

func getLastSeenMessageID(userID int64, channelID int64) (int64, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("HGET", lastSeenMessageIDKey(userID), channelID))
}
