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
	cnt, err := redis.Int64(conn.Do("ZCARD", messageIDsKey(chanID)))
	if err == redis.ErrNil {
		return 0, nil
	} else if err == nil {
		return cnt, nil
	}
	return 0, nil
}

func countUnreadMessages(chanID, lastSeenMsgID int64) (int64, error) {
	conn := pool.Get()
	defer conn.Close()
	cnt, err := redis.Int64(conn.Do("ZREVRANK", messageIDsKey(chanID), lastSeenMsgID))
	if err == redis.ErrNil {
		return 0, nil
	} else if err == nil {
		return cnt, nil
	}
	return 0, nil
}
