package commands

import (
	"github.com/TimIrwing/nyashka-butler/internal/bot/keyboard/pages"
	"github.com/TimIrwing/nyashka-butler/internal/types"
)

var handlers = map[string]func(CommandOptions) *types.Response{
	"config": func(opts CommandOptions) *types.Response {
		old := opts.Message
		m := old.New("Group settings:")
		m.SetKeyboardPage(pages.Config)
		return &types.Response{Message: m}
	},
}
