package commands

import "github.com/TimIrwing/nyashka-butler/internal/types"

var handlers = map[string]func(CommandOptions) *types.Response{
	"config": func(opts CommandOptions) *types.Response {
		old := opts.Message
		m := old.New("TODO")
		m.SetReplyID(old.GetID())
		return &types.Response{Message: m}
	},
}
