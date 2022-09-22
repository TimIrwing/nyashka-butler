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
		log.Printf("No such command: /%s", cmd)
		return nil
	}
	return handler(opts)
}
