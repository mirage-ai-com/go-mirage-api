// go-mirage-api
//
// Copyright 2023, Valerian Saliou
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

  // Categorize provided conversations
  data, err := client.Task.CategorizeConversations(mirage.RequestContext{}, mirage.CategorizeConversationsRequest {
    Conversations: []mirage.CategorizeConversationsRequestConversation {
      mirage.CategorizeConversationsRequestConversation {
        Transcript: []mirage.CategorizeConversationsRequestConversationTranscript {
          mirage.CategorizeConversationsRequestConversationTranscript {
            From: "customer",
            Text: "Hello! I have a question about the Crisp chatbot, I am trying to setup a week-end auto-responder, how can I do that?",
          },

          mirage.CategorizeConversationsRequestConversationTranscript {
            From: "agent",
            Text: "Hi. Baptiste here. I can provide you an example bot scenario that does just that if you'd like?",
          },
        },
      },
    },
  })

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Printf("Categorized conversations (raw): %s\n", data)
  }
}
