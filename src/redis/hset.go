package redis

import "fmt"

func handleHSet(db *DB, cmd Command) {

	if len(cmd.Args) != 3 {
		cmd.Result <- Response{Err: fmt.Errorf("HSET requires key field value")}
		return
	}

	key := cmd.Args[0]
	field := cmd.Args[1]
	value := cmd.Args[2]

	entry, ok := db.Get(key)

	var hash map[string]string
	var expiresAt int64 = 0

	if !ok {
		hash = make(map[string]string)
	} else {
		if entry.Type != HashType {
			cmd.Result <- Response{Err: fmt.Errorf("wrong type")}
			return
		}
		hash = entry.Value.(map[string]string)

		expiresAt = entry.ExpiresAt
	}

	hash[field] = value

	db.Set(key, hash, expiresAt, HashType)

	cmd.Result <- Response{Data: "OK"}
}
