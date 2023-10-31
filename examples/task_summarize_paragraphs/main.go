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

  // Summarize provided paragraphs
  data, err := client.Task.SummarizeParagraphs(mirage.SummarizeParagraphsRequest {
    Transcript: []mirage.SummarizeParagraphsRequestTranscript {
      mirage.SummarizeParagraphsRequestTranscript {
        Text: "GPT-4 is getting worse over time, not better.",
      },

      mirage.SummarizeParagraphsRequestTranscript {
        Text: "Many people have reported noticing a significant degradation in the quality of the model responses, but so far, it was all anecdotal.",
      },
    },
  })

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Printf("Summarized paragraphs (raw): %s\n", data)
  }
}
