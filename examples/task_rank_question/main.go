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

  // Generate rank data
  rankSource := "helpdesk"

  // Rank provided question results
  data, err := client.Task.RankQuestion(mirage.RankQuestionRequest {
    Question: "Hi! I am having issues setting up DNS records for my Crisp helpdesk. Can you help?",

    Context: mirage.RankQuestionRequestContext {
      Source: &rankSource,
      PrimaryID: "cf4ccdb5-df44-4668-a9e7-3ab31bebf89b",
    },
  })

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Printf("Ranked results (raw): %s\n", data)
  }
}
