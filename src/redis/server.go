package redis

import "fmt"

func EventLoop(queue <-chan Command) {
	db := NewDB()

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
