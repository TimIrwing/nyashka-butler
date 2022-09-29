package commands

import (
	"github.com/TimIrwing/nyashka-butler/internal/interfaces"
	"log"
)

type CommandOptions struct {
	Args    []string
	Message interfaces.Message
}

func Handle(cmd string, opts CommandOptions) interfaces.Message {
	handler, ok := handlers[cmd]
	if !ok {
		log.Printf("No such command: /%s", cmd)
		return nil
	}
	return handler(opts)
}
