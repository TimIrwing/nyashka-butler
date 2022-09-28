package types

type KeyboardPage string

type TextInteractionName string

type BotInfo struct {
	ID       int64
	UserName string
}

type User struct {
	ID int64
	// optional
	UserName string
}

type SettingsGeneral struct {
	IsGeneral bool
	// Super admin mode, executes admin orders only if they have enough permissions TODO which permissions exactly
	Super bool
}

type SettingsTextInteraction struct {
	TextInteractionName string
	Enabled             bool
	Reply               bool
	Randomize           bool
	Selected            int
	Text                []string
}
