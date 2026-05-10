package redis

import "fmt"

func handleSet(db *DB, cmd Command) {
	if len(cmd.Args) != 2 {
		cmd.Result <- Response{
			Err: fmt.Errorf("SET requires key value"),
		}
		return
	}

	key := cmd.Args[0]
	value := cmd.Args[1]

	db.data[key] = Entry{
		Type:  StringType,
		Value: value,
	}

	cmd.Result <- Response{
		Data: "OK",
	}
}
