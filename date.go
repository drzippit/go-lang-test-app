package main

import (
  "fmt"
  "context"


  "github.com/go-telegram/bot"
  "github.com/go-telegram/bot/models"
  "github.com/golang-module/carbon/v2"
)

func dateHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	reply := bot.EscapeMarkdown(carbon.Now().ToDateString())
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
