# go-mirage-api

[![Test and Build](https://github.com/mirage-ai-com/go-mirage-api/workflows/Test%20and%20Build/badge.svg)](https://github.com/mirage-ai-com/go-mirage-api/actions?query=workflow%3A%22Test+and+Build%22)

The Mirage API Golang wrapper. Access AI inference services.

Copyright 2023 Crisp IM SAS. See LICENSE for copying information.

* **📝 Implements**: [API Reference (V1)](https://docs.mirage-ai.com/references/api/v1/) at revision: 07/01/2025
* **😘 Maintainer**: [@valeriansaliou](https://github.com/valeriansaliou)

## Usage

Install the library:

```bash
go get github.com/mirage-ai-com/go-mirage-api@v1
```

Import the library:

```go
import "github.com/mirage-ai-com/go-mirage-api"
```

Construct a new authenticated Mirage client with your `user_id` and `secret_key` tokens.

```go
client := mirage.New("ui_xxxxxx", "sk_xxxxxx")
```

Then, consume the client eg. to transcribe a audio file containing speech to text:

```go
data, err := client.Task.TranscribeSpeech(mirage.RequestContext{}, mirage.TranscribeSpeechRequest {
  Locale: mirage.TranscribeSpeechRequestLocale {
    To: "en",
  },

  Media: mirage.TranscribeSpeechRequestMedia {
    Type: "audio/webm",
    URL: "https://files.mirage-ai.com/dash/terminal/samples/transcribe-speech/hey-there.weba",
  },
})
```

## Authentication

To authenticate against the API, get your tokens (`user_id` and `secret_key`).

Then, pass those tokens **once** when you instanciate the Mirage client as following:

```go
# Make sure to replace 'user_id' and 'secret_key' with your tokens
client = mirage.New("user_id", "secret_key")
```

## Resource Methods

This library implements all methods the Mirage API provides. See the [API docs](https://docs.mirage-ai.com/references/api/v1/) for a reference of available methods, as well as how returned data is formatted.

### Task API

#### ➡️ Transcribe Speech

* **Method:** `client.Task.TranscribeSpeech(ctx, data)`
* **Reference:** [Transcribe Speech](https://docs.mirage-ai.com/references/api/v1/#transcribe-speech)

* **Request:**

```go
client.Task.TranscribeSpeech(mirage.RequestContext{}, mirage.TranscribeSpeechRequest {
  Locale: mirage.TranscribeSpeechRequestLocale {
    To: "en",
  },

  Media: mirage.TranscribeSpeechRequestMedia {
    Type: "audio/webm",
    URL: "https://files.mirage-ai.com/dash/terminal/samples/transcribe-speech/hey-there.weba",
  },
})
```

* **Response:**

```json
{
  "reason": "processed",

  "data": {
    "locale": "en",

    "parts": [
      {
        "start": 5.0,
        "end": 9.0,
        "text": " I'm just speaking some seconds to see if the translation is correct"
      }
    ]
  }
}
```

#### ➡️ Answer Prompt

* **Method:** `client.Task.AnswerPrompt(ctx, data)`
* **Reference:** [Answer Prompt](https://docs.mirage-ai.com/references/api/v1/#answer-prompt)

* **Request:**

```go
client.Task.AnswerPrompt(mirage.RequestContext{}, mirage.AnswerPromptRequest {
  Prompt: "Generate an article about Alpacas",
})
```

* **Response:**

```json
{
  "reason": "processed",

  "data": {
    "answer": "The alpaca (Lama pacos) is a species of South American camelid mammal. It is similar to, and often confused with, the llama. However, alpacas are often noticeably smaller than llamas. The two animals are closely related and can successfully crossbreed. Both species are believed to have been domesticated from their wild relatives, the vicuña and guanaco. There are two breeds of alpaca: the Suri alpaca and the Huacaya alpaca.",
    "model": "medium"
  }
}
```

#### ➡️ Answer Question

* **Method:** `client.Task.AnswerQuestion(ctx, data)`
* **Reference:** [Answer Question](https://docs.mirage-ai.com/references/api/v1/#answer-question)

* **Request:**

```go
answerStart := "Sure,"

client.Task.AnswerQuestion(mirage.RequestContext{}, mirage.AnswerQuestionRequest {
  Question: "Should I pay more for that?",

  Answer: &mirage.AnswerQuestionRequestAnswer {
    Start: &answerStart,
  },

  Context: mirage.AnswerQuestionRequestContext {
    PrimaryID: "cf4ccdb5-df44-4668-a9e7-3ab31bebf89b",

    Conversation: mirage.AnswerQuestionRequestContextConversation {
      Messages: []mirage.AnswerQuestionRequestContextConversationMessage {
        mirage.AnswerQuestionRequestContextConversationMessage {
          From: "customer",
          Text: "Hey there!",
        },

        mirage.AnswerQuestionRequestContextConversationMessage {
          From: "agent",
          Text: "Hi. How can I help?",
        },

        mirage.AnswerQuestionRequestContextConversationMessage {
          From: "customer",
          Text: "I want to add more sub-domains to my website.",
        },
      },
    },
  },
})
```

* **Response:**

```json
{
  "reason": "processed",

  "data": {
    "answer": "You can add the Crisp chatbox to your website by following this guide: https://help.crisp.chat/en/article/how-to-add-crisp-chatbox-to-your-website-dkrg1d/ :)",
    "model": "medium",
    "sources": []
  }
}
```

#### ➡️ Answer Chat

* **Method:** `client.Task.AnswerChat(ctx, data)`
* **Reference:** [Answer Chat](https://docs.mirage-ai.com/references/api/v1/#answer-chat)

* **Request:**

```go
client.Task.AnswerChat(mirage.RequestContext{}, mirage.AnswerChatRequest {
  Context: mirage.AnswerChatRequestContext {
    Conversation: mirage.AnswerChatRequestContextConversation {
      Messages: []mirage.AnswerChatRequestContextConversationMessage {
        mirage.AnswerChatRequestContextConversationMessage {
          From: "customer",
          Text: "Hey there!",
        },

        mirage.AnswerChatRequestContextConversationMessage {
          From: "agent",
          Text: "Hi. How can I help?",
        },

        mirage.AnswerChatRequestContextConversationMessage {
          From: "customer",
          Text: "What is the weather in Nantes, France?",
        },
      },
    },
  },
  
  Tools: []mirage.AnswerChatRequestTool{
    mirage.AnswerChatRequestTool {
      Type: "function",

      Function: mirage.AnswerChatRequestToolFunction {
        Name: "get_current_weather",
        Description: "Get the current weather for a city",

        // Add parameters as JSON schema
        Parameters: nil,
      },
    },
  },
})
```

* **Response:**

```json
{
  "answer": "",
  "model": "medium",

  "tool_calls": [
    {
      "function": {
        "name": "get_current_weather",

        "arguments": {
          "city": "Nantes"
        }
      }
    }
  ]
}
```

#### ➡️ Summarize Paragraphs

* **Method:** `client.Task.SummarizeParagraphs(ctx, data)`
* **Reference:** [Summarize Paragraphs](https://docs.mirage-ai.com/references/api/v1/#summarize-paragraphs)

* **Request:**

```go
client.Task.SummarizeParagraphs(mirage.RequestContext{}, mirage.SummarizeParagraphsRequest {
  Paragraphs: []mirage.SummarizeParagraphsRequestParagraph {
    mirage.SummarizeParagraphsRequestParagraph {
      Text: "GPT-4 is getting worse over time, not better.",
    },

    mirage.SummarizeParagraphsRequestParagraph {
      Text: "Many people have reported noticing a significant degradation in the quality of the model responses, but so far, it was all anecdotal.",
    },
  },
})
```

* **Response:**

```json
{
  "reason": "processed",

  "data": {
    "summary": "GPT-4 is getting worse over time, not better. We have a new version of GPT-4 that is not improving, but it is regressing."
  }
}
```

#### ➡️ Summarize Conversation

* **Method:** `client.Task.SummarizeConversation(ctx, data)`
* **Reference:** [Summarize Conversation](https://docs.mirage-ai.com/references/api/v1/#summarize-conversation)

* **Request:**

```go
client.Task.SummarizeConversation(mirage.RequestContext{}, mirage.SummarizeConversationRequest {
  Transcript: []mirage.SummarizeConversationRequestTranscript {
    mirage.SummarizeConversationRequestTranscript {
      Name: "Valerian",
      Text: "Hello! I have a question about the Crisp chatbot, I am trying to setup a week-end auto-responder, how can I do that?",
    },

    mirage.SummarizeConversationRequestTranscript {
      Name: "Baptiste",
      Text: "Hi. Baptiste here. I can provide you an example bot scenario that does just that if you'd like?",
    },
  },
})
```

* **Response:**

```json
{
  "reason": "processed",

  "data": {
    "summary": "Valerian wants to set up a week-end auto-responder on Crisp chatbot. Baptiste can give him an example."
  }
}
```

#### ➡️ Categorize Conversations

* **Method:** `client.Task.CategorizeConversations(ctx, data)`
* **Reference:** [Categorize Conversations](https://docs.mirage-ai.com/references/api/v1/#categorize-conversations)

* **Request:**

```go
client.Task.CategorizeConversations(mirage.RequestContext{}, mirage.CategorizeConversationsRequest {
  Conversations: []mirage.CategorizeConversationsRequestConversation {
    mirage.CategorizeConversationsRequestConversation {
      Transcript: []mirage.CategorizeConversationsRequestConversationTranscript {
        mirage.CategorizeConversationsRequestConversationTranscript {
          From: "customer",
          Text: "Hello! I have a question about the Crisp chatbot, I am trying to setup a week-end auto-responder, how can I do that?",
        },

        mirage.CategorizeConversationsRequestConversationTranscript {
          From: "agent",
          Text: "Hi. Baptiste here. I can provide you an example bot scenario that does just that if you'd like?",
        },
      },
    },
  },
})
```

* **Response:**

```json
{
  "reason": "processed",

  "data": {
    "categories": [
      "Chatbot Configuration Issue"
    ]
  }
}
```

#### ➡️ Categorize Question

* **Method:** `client.Task.CategorizeQuestion(ctx, data)`
* **Reference:** [Categorize Question](https://docs.mirage-ai.com/references/api/v1/#categorize-question)

* **Request:**

```go
client.Task.CategorizeQuestion(mirage.RequestContext{}, mirage.CategorizeQuestionRequest {
  Question: "Hello. I have a question",
})
```

* **Response:**

```json
{
  "reason": "processed",

  "data": {
    "category": "opener"
  }
}
```

#### ➡️ Rank Question

* **Method:** `client.Task.RankQuestion(ctx, data)`
* **Reference:** [Rank Question](https://docs.mirage-ai.com/references/api/v1/#rank-question)

* **Request:**

```go
rankSource := "helpdesk"

client.Task.RankQuestion(mirage.RequestContext{}, mirage.RankQuestionRequest {
  Question: "Hi! I am having issues setting up DNS records for my Crisp helpdesk. Can you help?",

  Context: mirage.RankQuestionRequestContext {
    Source: &rankSource,
    PrimaryID: "cf4ccdb5-df44-4668-a9e7-3ab31bebf89b",
  },
})
```

* **Response:**

```json
{
  "reason": "processed",

  "data": {
    "results": [
      {
        "id": "15fd3f24-56c8-435e-af8e-c47d4cd6115c",
        "score": 9,
        "grouped_text": "Setup your Helpdesk domain name\ntutorials for most providers",

        "items": [
          {
            "source": "helpdesk",
            "primary_id": "51a32e4c-1cb5-47c9-bcc0-3e06f0dce90a",
            "secondary_id": "15fd3f24-56c8-435e-af8e-c47d4cd6115c",
            "text": "Setup your Helpdesk domain name\ntutorials for most providers",
            "timestamp": 1682002198552,

            "metadata": {
              "title": "Setup your Helpdesk domain name"
            }
          }
        ]
      }
    ]
  }
}
```

#### ➡️ Translate Text

* **Method:** `client.Task.TranslateText(ctx, data)`
* **Reference:** [Translate Text](https://docs.mirage-ai.com/references/api/v1/#translate-text)

* **Request:**

```go
translateType := "html"

client.Task.TranslateText(mirage.RequestContext{}, mirage.TranslateTextRequest {
  Locale: mirage.TranslateTextRequestLocale {
    From: "fr",
    To: "en",
  },

  Type: &translateType,
  Text: "Bonjour, comment puis-je vous aider <span translate=\"no\">Mr Saliou</span> ?",
})
```

* **Response:**

```json
{
  "reason": "processed",

  "data": {
    "translation": "Hi, how can I help you Mr Saliou?"
  }
}
```

#### ➡️ Fraud Spamicity

* **Method:** `client.Task.FraudSpamicity(ctx, data)`
* **Reference:** [Fraud Spamicity](https://docs.mirage-ai.com/references/api/v1/#fraud-spamicity)

* **Request:**

```go
client.Task.FraudSpamicity(mirage.RequestContext{}, mirage.FraudSpamicityRequest {
  Name: "Crisp",
  Domain: "crisp.chat",
  EmailDomain: "mail.crisp.chat",
})
```

* **Response:**

```json
{
  "reason": "processed",

  "data": {
    "fraud": false,
    "score": 0.13
  }
}
```

#### ➡️ Spam Conversation

* **Method:** `client.Task.SpamConversation(ctx, data)`
* **Reference:** [Spam Conversation](https://docs.mirage-ai.com/references/api/v1/#spam-conversation)

* **Request:**

```go
client.Task.SpamConversation(mirage.RequestContext{}, mirage.SpamConversationRequest {
  Sender: mirage.SpamConversationRequestSender {
    Name: "John Doe",
    Email: "john@example.com",
  },

  Transcript: []mirage.SpamConversationRequestTranscript {
    mirage.SpamConversationRequestTranscript {
      From: "customer",
      Origin: "chat",
      Text: "Hello, I would like to discuss your services",
    },
  },
})
```

* **Response:**

```json
{
  "reason": "processed",

  "data": {
    "class": "spam",
    "confidence": 0.93,
    "logprob": -0.10,

    "scores": {
      "gibberish": 0.0,
      "marketing": 0.45,
      "regular": 0.0,
      "spam": 0.93
    }
  }
}
```

#### ➡️ Spam Document

* **Method:** `client.Task.SpamDocument(ctx, data)`
* **Reference:** [Spam Document](https://docs.mirage-ai.com/references/api/v1/#spam-document)

* **Request:**

```go
client.Task.SpamDocument(mirage.RequestContext{}, mirage.SpamDocumentRequest {
  Name: "Spammy Domain",
  Domain: "spammy-domain.crisp.help",
  Title: "Spammy title",
  Content: "Spammy content",
})
```

* **Response:**

```json
{
  "reason": "processed",

  "data": {
    "class": "spam",
    "confidence": 0.82,
    "logprob": -0.10,

    "scores": {
      "gibberish": null,
      "marketing": null,
      "regular": 0.0,
      "spam": 0.82
    }
  }
}
```

### Data API

#### ➡️ Context Ingest

* **Method:** `client.Task.IngestContextData(ctx, data)`
* **Reference:** [Ingest Context Data](https://docs.mirage-ai.com/references/api/v1/#ingest-context-data)

* **Request:**

```go
ingestSecondaryID := "sec_6693a4a2-e33f-4cce-ba90-b7b5b0922c46"
ingestTertiaryID := "ter_de2bd6e7-74e1-440d-9a23-01964cd4b7da"
ingestText := "Text to index here..."
ingestSource := "chat"
ingestTimestamp := uint64(1682002198552)

ingestMetadata := map[string]string{
  "custom_key": "custom_value",
  "another_key": "another_value",
}

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
```

* **Response:**

```json
{
  "reason": "processed",

  "data": {
    "imported": true
  }
}
```
