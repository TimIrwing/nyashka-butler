package commands

import (
	"github.com/TimIrwing/nyashka-butler/internal/bot/keyboard/pages"
	"github.com/TimIrwing/nyashka-butler/internal/interfaces"
)

var handlers = map[string]func(CommandOptions) interfaces.Message{
	"config": func(opts CommandOptions) interfaces.Message {
		old := opts.Message
		m := old.New("Group settings:")
		m.SetKeyboardPage(pages.Config)
		return m
	},
}
