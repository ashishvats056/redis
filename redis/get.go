package redis

import "fmt"

func handleGet(db *DB, cmd Command) {
	if len(cmd.Args) != 1 {
		cmd.Result <- Response{
			Err: fmt.Errorf("GET requires key"),
		}
		return
	}

	key := cmd.Args[0]

	entry, ok := getEntry(db, key)
	if !ok {
		cmd.Result <- Response{
			Err: fmt.Errorf("key not found"),
		}
		return
	}

	if entry.Type != StringType {
		cmd.Result <- Response{
			Err: fmt.Errorf("wrong type"),
		}
		return
	}

	cmd.Result <- Response{
		Data: entry.Value.(string),
	}
}
