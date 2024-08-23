package main

import (
  "fmt"

  "github.com/golang-module/carbon/v2"
)

func printLog(resv string, send string, chatid int64) {
  logResv := carbon.Now().String() + " Message received: " + resv
  logSend := carbon.Now().String() + " Message sent: " + send
  logChatID := carbon.Now().String() + " Message received in chat: " + fmt.Sprint(chatid)

  fmt.Println(logResv)
  fmt.Println(logSend)
  fmt.Println(logChatID)
}
