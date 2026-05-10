package redis

import "fmt"

func handleLAdd(db *DB, cmd Command) {

	if len(cmd.Args) != 2 {
		cmd.Result <- Response{Err: fmt.Errorf("LADD requires key value")}
		return
	}

	key := cmd.Args[0]
	value := cmd.Args[1]

	entry, ok := db.Get(key)

	var list []string

	if !ok {
		list = []string{}
	} else {
		if entry.Type != ListType {
			cmd.Result <- Response{Err: fmt.Errorf("wrong type")}
			return
		}
		list = entry.Value.([]string)
	}

	list = append(list, value)

	db.Set(key, list, 0, ListType)

	cmd.Result <- Response{Data: "OK"}
}
