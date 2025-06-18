// go-mirage-api
//
// Copyright 2025, Valerian Saliou
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

  // Answer provided question
  data, err := client.Task.AnswerChat(mirage.RequestContext{}, mirage.AnswerChatRequest {
    Context: mirage.AnswerChatRequestContext {
      Conversation: mirage.AnswerChatRequestContextConversation {
        Messages: []mirage.AnswerChatRequestContextConversationMessage {
          mirage.AnswerChatRequestContextConversationMessage {
            From: "customer",
            Text: "Hey there!",
          },

          mirage.AnswerChatRequestContextConversationMessage {
            From: "agent",
            Text: "Hi. How can I help?",
          },

          mirage.AnswerChatRequestContextConversationMessage {
            From: "customer",
            Text: "What is the weather in Nantes, France?",
          },
        },
      },
    },
    
    Tools: []mirage.AnswerChatRequestTool{
      mirage.AnswerChatRequestTool {
        Type: "function",
        Function: mirage.AnswerChatRequestToolFunction {
          Name: "get_current_weather",
          Description: "Get the current weather for a city",

          // Add parameters as JSON schema
          Parameters: nil,
        },
      },
    },
  })

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Printf("Answer (raw): %s\n", data)
  }
}
