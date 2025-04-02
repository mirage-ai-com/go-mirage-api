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

  // Spam classification on email messages
  data, err := client.Task.SpamDocument(mirage.RequestContext{}, mirage.SpamDocumentRequest {
    Name: "Spammy Domain",
    Domain: "spammy-domain.crisp.help",
    Title: "Spammy title",
    Content: "Spammy content",
  })

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Printf("Spam document results (raw): %s\n", data)
  }
}
