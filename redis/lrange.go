package redis

import (
	"fmt"
	"strconv"
)

func handleLRange(db *DB, cmd Command) {

	if len(cmd.Args) != 3 {
		cmd.Result <- Response{Err: fmt.Errorf("LRANGE requires key start end")}
		return
	}

	key := cmd.Args[0]

	start, _ := strconv.Atoi(cmd.Args[1])
	end, _ := strconv.Atoi(cmd.Args[2])

	entry, ok := db.Get(key)
	if !ok {
		cmd.Result <- Response{Err: fmt.Errorf("key not found")}
		return
	}

	if entry.Type != ListType {
		cmd.Result <- Response{Err: fmt.Errorf("wrong type")}
		return
	}

	list := entry.Value.([]string)

	if start < 0 {
		start = 0
	}
	if end >= len(list) {
		end = len(list) - 1
	}

	if start > end {
		cmd.Result <- Response{Data: []string{}}
		return
	}

	cmd.Result <- Response{Data: list[start : end+1]}
}
