package commands

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"fmt"

)

func sendToLogsStart(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := b.SendMessage(-1001650787300, fmt.Sprintf("#START_COMMAND\nUser: <code>%s</code>\nUsername: @%s\nUserID: <code>%d</code>", ctx.EffectiveUser.Username, ctx.EffectiveUser.Username, ctx.EffectiveUser.Id), &gotgbot.SendMessageOpts{ParseMode: "html"})
	if err != nil {
		return err
	}
	return nil
}

func Start(b *gotgbot.Bot, ctx *ext.Context) error {
	sendToLogsStart(b, ctx)
	_, err := ctx.EffectiveMessage.Reply(b, "Hello There! I am a simple bot created in Go to generate complex passwords. My goal is to generate a password of the right length for you as quickly as possible so that you can use it wherever you need it. We do not log your passwords and our code is free and open source so you can clearly see that no surveillance is done on you. Below is a button with a link to Github, there is my source code\n\nUse /pass <length> to generate your password\nUse /developers to contact the developers",
		&gotgbot.SendMessageOpts{ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
				{Text: "Github", Url: "https://github.com/Yanninay/obiwanbot"},
			}},
		},
		})
	if err != nil {
		return err
	}
	
	return nil
}
