package redis

import (
	"fmt"
	"time"
)

func handleTTL(db *DB, cmd Command) {

	if len(cmd.Args) != 1 {
		cmd.Result <- Response{Err: fmt.Errorf("TTL requires key")}
		return
	}

	key := cmd.Args[0]

	entry, ok := db.Get(key)
	if !ok {
		cmd.Result <- Response{Data: -2}
		return
	}

	if entry.ExpiresAt == 0 {
		cmd.Result <- Response{Data: -1}
		return
	}

	remaining := time.Until(time.Unix(0, entry.ExpiresAt)).Seconds()

	cmd.Result <- Response{Data: int64(remaining)}
}
