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

  // Fraud check classification on scammy websites
  data, err := client.Task.FraudSpamicity(mirage.RequestContext{}, mirage.FraudSpamicityRequest {
    Name: "Crisp",
    Domain: "crisp.chat",
    EmailDomain: "mail.crisp.chat",
  })

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Printf("Fraud check spamicity results (raw): %s\n", data)
  }
}
