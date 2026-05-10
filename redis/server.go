package redis

import "fmt"

func EventLoop(queue <-chan Command) {
	db := DB{
		data: make(map[string]Entry),
	}

	startCleaner(&db)

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
