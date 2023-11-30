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

  // Generate answer data
  answerStart := "Sure,"

  // Answer provided question
  data, err := client.Task.AnswerQuestion(mirage.AnswerQuestionRequest {
    Question: "Should I pay more for that?",

    Answer: &mirage.AnswerQuestionRequestAnswer {
      Start: &answerStart,
    },

    Context: mirage.AnswerQuestionRequestContext {
      PrimaryID: "cf4ccdb5-df44-4668-a9e7-3ab31bebf89b",

      Conversation: mirage.AnswerQuestionRequestContextConversation {
        Messages: []mirage.AnswerQuestionRequestContextConversationMessage {
          mirage.AnswerQuestionRequestContextConversationMessage {
            From: "customer",
            Text: "Hey there!",
          },

          mirage.AnswerQuestionRequestContextConversationMessage {
            From: "agent",
            Text: "Hi. How can I help?",
          },

          mirage.AnswerQuestionRequestContextConversationMessage {
            From: "customer",
            Text: "I want to add more sub-domains to my website.",
          },
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
