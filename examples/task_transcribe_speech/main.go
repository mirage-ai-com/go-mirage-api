// go-mirage-api
//
// Copyright 2023, Valerian Saliou
// Author: Valerian Saliou <valerian@valeriansaliou.name>

package main

import (
  "github.com/mirage-ai-com/go-mirage-api/mirage"
  "fmt"
)

const (
  CONFIG_USER_ID = "ui_xxxxxx"
  CONFIG_SECRET_KEY = "sk_xxxxxx"
)

func main() {
  client := mirage.New(CONFIG_USER_ID, CONFIG_SECRET_KEY)

  // Transcribe provided audio file
  data, err := client.Task.TranscribeSpeech(mirage.TranscribeSpeechRequest {
	  Locale: mirage.TranscribeSpeechRequestLocale {
	    To: "en",
	  },

	  Media: mirage.TranscribeSpeechRequestMedia {
	    Type: "audio/webm",
	    URL: "https://files.mirage-ai.com/dash/terminal/samples/transcribe-speech/hey-there.weba",
	  },
	})

  if err != nil {
    fmt.Printf("Error: %s", err)
  } else {
    fmt.Printf("Transcription (raw): %s\n", data)
  }
}
