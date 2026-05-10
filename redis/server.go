package redis

import "fmt"

var (
	get    ValueType = StringType
	set    ValueType = StringType
	del    ValueType = StringType
	index  ValueType = StringType
	mget   ValueType = StringType
	hget   ValueType = HashType
	hset   ValueType = HashType
	ladd   ValueType = ListType
	lrange ValueType = ListType
)

func EventLoop(queue <-chan Command) {
	db := DB{
		data: make(map[string]Entry),
	}

	for cmd := range queue {
		handler, ok := commands[cmd.Name]

		if !ok {
			cmd.Result <- Response{
				Err: fmt.Errorf("unknown command"),
			}
			continue
		}

		handler(&db, cmd)
	}
}
