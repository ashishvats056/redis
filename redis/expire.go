package redis

import (
	"fmt"
	"strconv"
	"time"
)

func handleExpire(db *DB, cmd Command) {

	if len(cmd.Args) != 2 {
		cmd.Result <- Response{Err: fmt.Errorf("EXPIRE requires key seconds")}
		return
	}

	key := cmd.Args[0]
	seconds, err := strconv.Atoi(cmd.Args[1])
	if err != nil || seconds <= 0 {
		cmd.Result <- Response{Err: fmt.Errorf("invalid TTL")}
		return
	}

	entry, ok := db.Get(key)
	if !ok {
		cmd.Result <- Response{Data: 0}
		return
	}

	exp := time.Now().Add(time.Duration(seconds) * time.Second).UnixNano()

	db.Set(key, entry.Value, exp, entry.Type)

	cmd.Result <- Response{Data: 1}
}
