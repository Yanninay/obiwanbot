package commands

// generate a random password

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"math/rand"
	"strconv"
	"time"
)

func Main(b *gotgbot.Bot, ctx *ext.Context) error {
	rand.Seed(time.Now().UnixNano())
	return nil
}

func sendToLogs(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := b.SendMessage(-1001650787300, fmt.Sprintf("#PASS_COMMAND\nUser: <code>%s</code>\nUsername: @%s\nUserID: <code>%d</code>", ctx.EffectiveUser.Username, ctx.EffectiveUser.Username, ctx.EffectiveUser.Id), &gotgbot.SendMessageOpts{ParseMode: "html"})
	if err != nil {
		return err
	}
	return nil
}

func Pass(b *gotgbot.Bot, ctx *ext.Context) error {
	if len(ctx.Args()) == 1 {
		sendToLogs(b, ctx)
		_, err := ctx.EffectiveMessage.Reply(b, "You didn't mention the arguments!\nUsage: /pass <length>", &gotgbot.SendMessageOpts{})
		return err
	}
	passLength, err := strconv.ParseInt(ctx.Args()[1], 10, 64)
	if err != nil {
		sendToLogs(b, ctx)
		_, err := ctx.EffectiveMessage.Reply(b, "Unfortunately, the argument you mentioned is not a number. Please check it, and specify it again", &gotgbot.SendMessageOpts{})
		return err
	}
	if passLength < 4 || passLength > 100 {
		sendToLogs(b, ctx)
		_, err := ctx.EffectiveMessage.Reply(b, "The argument you mentioned is not in the range of 4 to 128. Please check it, and specify it again", &gotgbot.SendMessageOpts{})
		return err
	}
	_, _ = b.SendChatAction(ctx.EffectiveChat.Id, "typing", &gotgbot.SendChatActionOpts{})
	password := generatePassword(ctx.Args()[1])
	if len(password) > 98 {
		sendToLogs(b, ctx)
		_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf(`Here's your password: %s`, password), &gotgbot.SendMessageOpts{})
		return err
	} else {
		sendToLogs(b, ctx)
		password := generatePassword(ctx.Args()[1])
		_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf(`Here's your password: <code>%s</code>`, password), &gotgbot.SendMessageOpts{
			ParseMode: "html",
		})
		return err
	}
}

func generatePassword(passwordLength string) string {
	// convert string to int
	length, _ := strconv.ParseInt(passwordLength, 10, 64)
	password := randomString(int(length))
	return password
}

func randomString(n int) string {
	// create random string
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+{}|<>?-=[]\\;':,./~`"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
