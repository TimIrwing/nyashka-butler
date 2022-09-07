package bot

import "github.com/TimIrwing/nyashka-butler/internal/global"

func (u *Update) IsCommand() bool {
	m := u._u.Message
	if m == nil {
		return false
	}
	for _, e := range m.Entities {
		if e.Type == "bot_command" {
			return true
		}
	}
	return false
}

type Command struct {
	name   string
	handle func() // TODO
}

func (u *Update) HandleCommand(global *global.Global) {
	// TODO admin permissions
}
