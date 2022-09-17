package commands

import (
	"github.com/TimIrwing/nyashka-butler/internal/interfaces"
	"github.com/TimIrwing/nyashka-butler/internal/types"
	"log"
)

type CommandOptions struct {
	Args    []string
	Message interfaces.Message
}

func Trigger(cmd string, opts CommandOptions) *types.Response {
	handler, ok := handlers[cmd]
	if !ok {
		log.Println("No such command")
		return nil
	}
	return handler(opts)
}
