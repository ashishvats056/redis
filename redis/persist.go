package redis

import "fmt"

func handlePersist(db *DB, cmd Command) {

	if len(cmd.Args) != 1 {
		cmd.Result <- Response{Err: fmt.Errorf("PERSIST requires key")}
		return
	}

	key := cmd.Args[0]

	entry, ok := db.Get(key)
	if !ok {
		cmd.Result <- Response{Data: 0}
		return
	}

	db.Set(key, entry.Value, 0, entry.Type)

	cmd.Result <- Response{Data: 1}
}
