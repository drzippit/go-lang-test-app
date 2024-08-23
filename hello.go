package main

import (
  "context"
  "fmt"

  "github.com/go-telegram/bot"
  "github.com/go-telegram/bot/models"
)


func helloHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
  reply := "Hello, *" + bot.EscapeMarkdown(update.Message.From.FirstName) + "*"
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
