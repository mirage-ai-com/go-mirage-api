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
  data, err := client.Task.TranscribeSpeech({
	  "locale": {
	    "to": "en"
	  },

	  "media": {
	    "type": "audio/mp3",
	    "url": "https://storage.crisp.chat/users/upload/session/5acfdb5400c15c00/audio1681224631050_9elgef.mp3"
	  }
	})

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Printf("Transcription (raw): %s\n", data)
  }
}
