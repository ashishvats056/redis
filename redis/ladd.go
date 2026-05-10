package redis

import "fmt"

func handleLAdd(db *DB, cmd Command) {

	if len(cmd.Args) != 2 {
		cmd.Result <- Response{Err: fmt.Errorf("LADD requires key value")}
		return
	}

	key := cmd.Args[0]
	value := cmd.Args[1]

	entry, ok := getEntry(db, key)

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

	db.data[key] = Entry{
		Type:  ListType,
		Value: list,
	}

	cmd.Result <- Response{Data: "OK"}
}
