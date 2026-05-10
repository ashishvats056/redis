package redis

import (
	"fmt"
	"strconv"
	"time"
)

func handleSet(db *DB, cmd Command) {
	if len(cmd.Args) < 2 {
		cmd.Result <- Response{
			Err: fmt.Errorf("SET requires key value [ttl]"),
		}
		return
	}

	key := cmd.Args[0]
	value := cmd.Args[1]

	var expiresAt int64 = 0

	// optional TTL
	if len(cmd.Args) == 3 {
		ttlSeconds, err := strconv.Atoi(cmd.Args[2])
		if err == nil && ttlSeconds > 0 {
			expiresAt = time.Now().Add(time.Duration(ttlSeconds) * time.Second).UnixNano()
		}
	}

	db.data[key] = Entry{
		Type:      StringType,
		Value:     value,
		ExpiresAt: expiresAt,
	}

	cmd.Result <- Response{
		Data: "OK",
	}
}
