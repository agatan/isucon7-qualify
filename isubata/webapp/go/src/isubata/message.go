package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

func messagesKey(chanID int64) string {
	return fmt.Sprintf("messages:%d", chanID)
}

func messageIDsKey(chanID int64) string {
	return fmt.Sprintf("message_ids:%d", chanID)
}

func userIDByMessageID() string {
	return "user_id_by_message_id"
}

func addMessage(chanID, userID int64, content string, now int64) (int64, error) {
	conn := pool.Get()
	defer conn.Close()
	id := now
	_, err := conn.Do("ZADD", messagesKey(chanID), id, content)
	if err != nil {
		return 0, err
	}
	_, err = conn.Do("ZADD", messageIDsKey(chanID), id, id)
	if err != nil {
		return 0, err
	}

	_, err = conn.Do("HSET", userIDByMessageID(), id, userID)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func getMessagesByChannelID(chanID, lastID int64) ([]*Message, error) {
	conn := pool.Get()
	defer conn.Close()
	id := int64(-1)
	var err error
	if lastID != 0 {
		id, err = redis.Int64(conn.Do("ZREVRANK", messageIDsKey(chanID), lastID))
		if err != nil {
			return nil, err
		}
	}
	values, err := redis.Values(conn.Do("ZREVRANGE", messagesKey(chanID), 0, id, "WITHSCORES"))
	l := len(values) / 2
	if l == 0 || err == redis.ErrNil {
		return []*Message{}, nil
	}
	msgs := make([]*Message, l)
	for i := 0; i < l; i++ {
		msgs[i] = &Message{}
		score := ""
		values, err = redis.Scan(values, &msgs[i].Content, &score)
		if err != nil {
			return nil, err
		}
		msgs[i].ID = scoreToID(score)
		msgs[i].ChannelID = chanID
		msgs[i].CreatedAt = time.Unix(msgs[i].ID, 0)
		msgs[i].UserID, err = redis.Int64(conn.Do("HGET", userIDByMessageID(), msgs[i].ID))
		if err != nil {
			return nil, err
		}
	}
	return msgs, nil
}

func getMessagesWithPage(chanID, page, per int64) ([]*Message, error) {
	conn := pool.Get()
	defer conn.Close()
	offset := per * (page - 1)
	values, err := redis.Values(conn.Do("ZREVRANGE", messagesKey(chanID), offset, offset+per, "WITHSCORES"))
	l := len(values) / 2
	if l == 0 || err == redis.ErrNil {
		return []*Message{}, nil
	}
	msgs := make([]*Message, l)
	for i := 0; i < l; i++ {
		msgs[i] = &Message{}
		score := ""
		values, err = redis.Scan(values, &msgs[i].Content, &score)
		if err != nil {
			return nil, err
		}
		msgs[i].ID = scoreToID(score)
		msgs[i].ChannelID = chanID
		msgs[i].CreatedAt = time.Unix(msgs[i].ID, 0)
		msgs[i].UserID, err = redis.Int64(conn.Do("HGET", userIDByMessageID(), msgs[i].ID))
		if err != nil {
			return nil, err
		}
	}
	return msgs, nil
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

func scoreToID(score string) int64 {
	id, _ := strconv.Atoi(score)
	return int64(id)
	// pair := strings.Split(score, "e+")
	// p0, _ := strconv.ParseFloat(pair[0], 64)
	// p1, _ := strconv.Atoi(pair[1])
	// return int64(float64(p0) * math.Pow(10, float64(p1)))
}
