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

  // Generate context ingest data
  ingestSecondaryID := "sec_6693a4a2-e33f-4cce-ba90-b7b5b0922c46"
  ingestTertiaryID := "ter_de2bd6e7-74e1-440d-9a23-01964cd4b7da"
  ingestText := "Text to index here..."
  ingestSource := "chat"
  ingestTimestamp := uint64(1682002198552)

  ingestMetadata := map[string]string{
    "custom_key": "custom_value",
    "another_key": "another_value",
  }

  // Ingest provided context data
  data, err := client.Task.IngestContextData(mirage.RequestContext{}, mirage.IngestContextDataRequest {
    Items: []mirage.IngestContextDataRequestItem {
      mirage.IngestContextDataRequestItem {
        Operation: "index",
        PrimaryID: "pri_cf44dd72-4ba9-4754-8fb3-83c4261243c4",
        SecondaryID: &ingestSecondaryID,
        TertiaryID: &ingestTertiaryID,

        Text: &ingestText,
        Source: &ingestSource,
        Timestamp: &ingestTimestamp,

        Metadata: &ingestMetadata,
      },
    },
  })

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Printf("Ingest context data (raw): %s\n", data)
  }
}
