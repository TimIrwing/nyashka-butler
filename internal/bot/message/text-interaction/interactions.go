package text_interaction

import (
	"fmt"
	"github.com/TimIrwing/nyashka-butler/internal/interfaces"
	"github.com/TimIrwing/nyashka-butler/internal/types"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
)

const TextNotSetFormat = "Interaction text is not set, somebody probably should change some settings. Idk set interaction text or turn off the feature (%s)"

func GetName(m *tgbotapi.Message, botInfo *types.BotInfo) types.TextInteractionName {
	reply := m.ReplyToMessage
	switch {
	case len(m.NewChatMembers) > 0:
		return Welcome
	case m.LeftChatMember != nil:
		return Goodbye
	case reply != nil && reply.From.ID == botInfo.ID:
		return ReplyToBot
	default:
		for _, e := range m.Entities {
			if e.IsMention() && m.Text[e.Offset+1:e.Offset+e.Length] == botInfo.UserName {
				return BotMention
			}
		}
	}
	return NoInteraction
}

func Handle(msg interfaces.Message, settings interfaces.Settings) interfaces.Message {
	name := msg.GetTextInteraction()
	if name == NoInteraction {
		return nil
	}

	config := settings.GetTextInteraction(name)
	if !config.Enabled {
		return nil
	}

	count := len(config.Text)
	var i int
	if config.Randomize && count > 0 {
		i = rand.Intn(count)
	} else {
		i = config.Selected
	}

	var text string
	if count == 0 {
		text = fmt.Sprintf(TextNotSetFormat, name)
	} else if i >= count {
		text = config.Text[count-1]
	} else {
		text = config.Text[i]
	}

	res := msg.New(text)
	if config.Reply {
		res.SetReplyID(msg.GetID())
	}
	return res
}
