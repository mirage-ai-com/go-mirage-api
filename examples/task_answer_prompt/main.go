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

  // Answer provided prompt
  data, err := client.Task.AnswerPrompt(mirage.AnswerPromptRequest {
    Prompt: "Generate an article about Alpacas",
  })

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Printf("Answer (raw): %s\n", data)
  }
}
