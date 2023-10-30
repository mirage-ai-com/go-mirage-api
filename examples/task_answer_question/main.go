// go-mirage-api
//
// Copyright 2023, Valerian Saliou
// Author: Valerian Saliou <valerian@valeriansaliou.name>

package main

import (
  "github.com/mirage-ai-com/go-mirage-api/mirage/v1"
  "fmt"
)

const (
  CONFIG_USER_ID = "ui_xxxxxx"
  CONFIG_SECRET_KEY = "sk_xxxxxx"
)

func main() {
  client := mirage.New(CONFIG_USER_ID, CONFIG_SECRET_KEY)

  // Answer provided question
  data, err := client.Task.AnswerQuestion({
    "question": "Should I pay more for that?",

    "answer": {
      "start": "Sure,"
    },

    "context": {
      "team": {
        "id": "cf4ccdb5-df44-4668-a9e7-3ab31bebf89b",
        "name": "Crisp"
      },

      "transcripts": {
        "conversation": {
          "messages": [
            {
              "from": "customer",
              "text": "Hey there!"
            },

            {
              "from": "agent",
              "text": "Hi. How can I help?"
            },

            {
              "from": "customer",
              "text": "I want to add more sub-domains to my website."
            }
          ]
        },

        "related": [
          {
            "messages": [
              {
                "from": "customer",
                "text": "Hi, does the \"per website\" pricing include sub-domains?"
              },

              {
                "from": "agent",
                "text": "Hi, yes, it includes sub-domains"
              },

              {
                "from": "customer",
                "text": "Perfect thanks!"
              }
            ]
          }
        ]
      }
    }
  })

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Printf("Answer (raw): %s\n", data)
  }
}
