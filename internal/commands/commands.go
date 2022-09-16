package commands

import (
    "log"
)

type Message interface {
    New(string) Message
    SetReply(int)
    GetID() int
    Send()
}

type CommandOptions struct {
    Args    []string
    Message Message
}

var commands = map[string]func(CommandOptions){
    "config": func(opts CommandOptions) {
        old := opts.Message
        m := old.New("TODO")
        m.SetReply(old.GetID())
        m.Send()
    },
}

func Trigger(cmd string, opts CommandOptions) {
    handler, ok := commands[cmd]
    if !ok {
        log.Println("No such command")
        return
    }
    handler(opts)
}
