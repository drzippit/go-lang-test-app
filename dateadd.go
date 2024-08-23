package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/golang-module/carbon/v2"
)

func dateaddHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	str := update.Message.Text

	parts := strings.Split(str, " ")
	if len(parts) > 1 {
		number, err := strconv.Atoi(parts[1])
		if err != nil {
			reply := "Error converting string to int"
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:    update.Message.Chat.ID,
				Text:      reply,
				ParseMode: models.ParseModeMarkdown,
			})
			printLog(update.Message.Text, reply, update.Message.Chat.ID)
			if err != nil {
				fmt.Printf("Failed to send message: %v", err)
			}
			return
		}

		reply := carbon.Now().AddDays(number).ToDateString()
		_, err = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   reply,
		})
		printLog(update.Message.Text, reply, update.Message.Chat.ID)
		if err != nil {
			fmt.Printf("Failed to send message: %v", err)
		}

	} else {

		reply := "Please provide a number to add to the date"
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      reply,
			ParseMode: models.ParseModeMarkdown,
		})
		printLog(update.Message.Text, reply, update.Message.Chat.ID)
		if err != nil {
			fmt.Printf("Failed to send message: %v", err)
		}
	}
}
