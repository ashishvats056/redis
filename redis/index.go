package redis

import (
	"fmt"
	"strconv"
)

func handleIndex(db *DB, cmd Command) {

	if len(cmd.Args) != 2 {
		cmd.Result <- Response{Err: fmt.Errorf("INDEX requires key index")}
		return
	}

	key := cmd.Args[0]

	idx, err := strconv.Atoi(cmd.Args[1])
	if err != nil {
		cmd.Result <- Response{Err: fmt.Errorf("invalid index")}
		return
	}

	entry, ok := getEntry(db, key)
	if !ok {
		cmd.Result <- Response{Err: fmt.Errorf("key not found")}
		return
	}

	if entry.Type != ListType {
		cmd.Result <- Response{Err: fmt.Errorf("wrong type")}
		return
	}

	list := entry.Value.([]string)

	if idx < 0 || idx >= len(list) {
		cmd.Result <- Response{Err: fmt.Errorf("index out of range")}
		return
	}

	cmd.Result <- Response{Data: list[idx]}
}
