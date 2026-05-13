package redis

import "fmt"

func handleDel(db *DB, cmd Command) {
	if len(cmd.Args) != 1 {
		cmd.Result <- Response{
			Err: fmt.Errorf("DEL requires key"),
		}
		return
	}

	key := cmd.Args[0]

	if _, exists := db.Get(key); !exists {
		cmd.Result <- Response{
			Err: fmt.Errorf("key not found"),
		}
		return
	}

	db.Delete(key)

	cmd.Result <- Response{
		Data: "OK",
	}
}
