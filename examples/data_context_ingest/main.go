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

  // Ingest provided context data
  data, err := client.Task.IngestContextData(mirage.IngestContextDataRequest {
    Items: []mirage.IngestContextDataRequestItem {
      mirage.IngestContextDataRequestItem {
        Operation: "index",
        PrimaryID: "pri_cf44dd72-4ba9-4754-8fb3-83c4261243c4",
        SecondaryID: "sec_6693a4a2-e33f-4cce-ba90-b7b5b0922c46",
        TertiaryID: "ter_de2bd6e7-74e1-440d-9a23-01964cd4b7da",

        Text: "Text to index here...",
        Source: "chat",
        Timestamp: 1682002198552,

        Metadata: {
          CustomKey: "custom_value",
          AnotherKey: "another_value",
        },
      },
    },
  })

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Printf("Ingest context data (raw): %s\n", data)
  }
}
