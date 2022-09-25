package settings

type General struct {
	IsGeneral bool
	// Super admin mode, executes admin orders only if they have enough permissions TODO which permissions exactly
	Super bool
}

type TextInteraction struct {
	TextInteractionName string
	Enabled             bool
	Reply               bool
	Randomize           bool
	Selected            int
	Text                []string
}
