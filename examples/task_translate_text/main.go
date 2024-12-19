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

  // Generate translate data
  translateType := "html"

  // Translate provided text
  data, err := client.Task.TranslateText(mirage.RequestContext{}, mirage.TranslateTextRequest {
    Locale: mirage.TranslateTextRequestLocale {
      From: "fr",
      To: "en",
    },

    Type: &translateType,
    Text: "Bonjour, comment puis-je vous aider <span translate=\"no\">Mr Saliou</span> ?",
	})

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Printf("Translation (raw): %s\n", data)
  }
}
