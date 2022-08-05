package commands

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"fmt"

)

func sendToLogsDev(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := b.SendMessage(-1001650787300, fmt.Sprintf("#DEVELOPERS_COMMAND\nUser: <code>%s</code>\nUsername: @%s\nUserID: <code>%d</code>", ctx.EffectiveUser.Username, ctx.EffectiveUser.Username, ctx.EffectiveUser.Id), &gotgbot.SendMessageOpts{ParseMode: "html"})
	if err != nil {
		return err
	}
	return nil
}

func Developers(b *gotgbot.Bot, ctx *ext.Context) error {
	sendToLogsDev(b, ctx)
	_, err := ctx.EffectiveMessage.Reply(b, "Thank you for using our bot! However, if you need support, here are the developers' contacts you can contact at any time",
		&gotgbot.SendMessageOpts{ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
				{Text: "Majes", Url: "https://t.me/mrmajesripper"},
				{Text: "~$ man dwm", Url: "https://t.me/GNUarch30"},
			}},
		},
		})
	if err != nil {
		return err
	}
	
	return nil
}
