package redis

func handleMGet(db *DB, cmd Command) {

	result := make([]string, 0, len(cmd.Args))

	for _, key := range cmd.Args {

		entry, ok := db.data[key]

		if !ok || entry.Type != StringType {
			result = append(result, "")
			continue
		}

		result = append(result, entry.Value.(string))
	}

	cmd.Result <- Response{
		Data: result,
	}
}
