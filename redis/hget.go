package redis

import "fmt"

func handleHGet(db *DB, cmd Command) {

	if len(cmd.Args) != 2 {
		cmd.Result <- Response{Err: fmt.Errorf("HGET requires key field")}
		return
	}

	key := cmd.Args[0]
	field := cmd.Args[1]

	entry, ok := getEntry(db, key)
	if !ok {
		cmd.Result <- Response{Err: fmt.Errorf("key not found")}
		return
	}

	if entry.Type != HashType {
		cmd.Result <- Response{Err: fmt.Errorf("wrong type")}
		return
	}

	hash := entry.Value.(map[string]string)

	val, ok := hash[field]
	if !ok {
		cmd.Result <- Response{Err: fmt.Errorf("field not found")}
		return
	}

	cmd.Result <- Response{Data: val}
}
