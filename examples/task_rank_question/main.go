// go-mirage-api
//
// Copyright 2023, Valerian Saliou
// Author: Valerian Saliou <valerian@valeriansaliou.name>

package main

import (
  "github.com/mirage-ai-com/go-mirage-api/mirage/v3"
  "fmt"
)

const (
  CONFIG_USER_ID = "ui_xxxxxx"
  CONFIG_SECRET_KEY = "sk_xxxxxx"
)

func main() {
  client := mirage.New(CONFIG_USER_ID, CONFIG_SECRET_KEY)

  // Rank provided question results
  data, err := client.Task.RankQuestion(mirage.RankQuestionRequest {
    Question: "Hi! I am having issues setting up DNS records for my Crisp helpdesk. Can you help?",
    Source: "helpdesk",

    Context: mirage.RankQuestionRequestContext {
      Team: mirage.RankQuestionRequestContextTeam {
        ID: "cf4ccdb5-df44-4668-a9e7-3ab31bebf89b",
        Name: "Crisp",
      },
    },
  })

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Printf("Ranked results (raw): %s\n", data)
  }
}