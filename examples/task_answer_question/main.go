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
      Team: mirage.AnswerQuestionRequestContextTeam {
        ID: "cf4ccdb5-df44-4668-a9e7-3ab31bebf89b",
        Name: "Crisp",
      },

      Transcripts: mirage.AnswerQuestionRequestContextTranscripts {
        Conversation: mirage.AnswerQuestionRequestContextTranscriptsConversation {
          Messages: []mirage.AnswerQuestionRequestContextTranscript {
            mirage.AnswerQuestionRequestContextTranscript {
              From: "customer",
              Text: "Hey there!",
            },

            mirage.AnswerQuestionRequestContextTranscript {
              From: "agent",
              Text: "Hi. How can I help?",
            },

            mirage.AnswerQuestionRequestContextTranscript {
              From: "customer",
              Text: "I want to add more sub-domains to my website.",
            },
          },
        },

        Related: &[]mirage.AnswerQuestionRequestContextTranscriptsRelated {
          mirage.AnswerQuestionRequestContextTranscriptsRelated {
            Messages: []mirage.AnswerQuestionRequestContextTranscript {
              mirage.AnswerQuestionRequestContextTranscript {
                From: "customer",
                Text: "Hi, does the \"per website\" pricing include sub-domains?",
              },

              mirage.AnswerQuestionRequestContextTranscript {
                From: "agent",
                Text: "Hi, yes, it includes sub-domains",
              },

              mirage.AnswerQuestionRequestContextTranscript {
                From: "customer",
                Text: "Perfect thanks!",
              },
            },
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
