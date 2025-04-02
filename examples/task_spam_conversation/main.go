// go-mirage-api
//
// Copyright 2024, Valerian Saliou
// Author: Valerian Saliou <valerian@valeriansaliou.name>

package main

import (
  "github.com/mirage-ai-com/go-mirage-api"
  "fmt"
)

const (
  CONFIG_USER_ID = "ui_xxxxxx"
  CONFIG_SECRET_KEY = "sk_xxxxxx"
)

func main() {
  client := mirage.New(CONFIG_USER_ID, CONFIG_SECRET_KEY)

  // Spam classification on email messages
  data, err := client.Task.SpamConversation(mirage.RequestContext{}, mirage.SpamConversationRequest {
    Sender: mirage.SpamConversationRequestSender {
      Name: "John Doe",
      Email: "john@example.com",
    },

    Transcript: []mirage.SpamConversationRequestTranscript {
      mirage.SpamConversationRequestTranscript {
        From: "customer",
        Origin: "chat",
        Text: "Hello, I would like to discuss your services",
      },
    },
  })

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Printf("Spam conversation results (raw): %s\n", data)
  }
}
