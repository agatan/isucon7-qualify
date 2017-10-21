package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func messageIDsKey(chanID int64) string {
	return fmt.Sprintf("message_ids:%d", chanID)
}

func appendMessageID(m *Message) error {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("ZADD", messageIDsKey(m.ChannelID), m.CreatedAt.UnixNano(), m.ID)
	return err
}

func countMessages(chanID int64) (int64, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("ZCARD", messageIDsKey(chanID)))
}

func countUnreadMessages(chanID, lastSeenMsgID int64) (int64, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("ZREVRANK", messageIDsKey(chanID), lastSeenMsgID))
}
