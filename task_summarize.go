// go-mirage-api
//
// Copyright 2023, Valerian Saliou
// Author: Valerian Saliou <valerian@valeriansaliou.name>

package mirage


// SummarizeParagraphsRequest mapping
type SummarizeParagraphsRequest struct {
  *Locale     SummarizeParagraphsRequestLocale       `json:"locale,omitempty"`
  Paragraphs  []SummarizeParagraphsRequestParagraph  `json:"paragraphs"`
}

// SummarizeParagraphsRequestLocale mapping
type SummarizeParagraphsRequestLocale struct {
  To  string  `json:"to"`
}

// SummarizeParagraphsRequestParagraph mapping
type SummarizeParagraphsRequestParagraph struct {
  Text  string  `json:"text"`
}

// SummarizeConversationRequest mapping
type SummarizeConversationRequest struct {
  *Locale     SummarizeConversationRequestLocale        `json:"locale,omitempty"`
  Transcript  []SummarizeConversationRequestTranscript  `json:"transcript"`
}

// SummarizeConversationRequestLocale mapping
type SummarizeConversationRequestLocale struct {
  To  string  `json:"to"`
}

// SummarizeConversationRequestTranscript mapping
type SummarizeConversationRequestTranscript struct {
  Name  string  `json:"name"`
  Text  string  `json:"text"`
}


// SummarizeParagraphsResponseData mapping
type SummarizeParagraphsResponseData struct {
  Data  *SummarizeParagraphsResponse  `json:"data"`
}

// SummarizeParagraphsResponse mapping
type SummarizeParagraphsResponse struct {
  Summary  string  `json:"summary"`
}

// SummarizeConversationResponseData mapping
type SummarizeConversationResponseData struct {
  Data  *SummarizeConversationResponse  `json:"data"`
}

// SummarizeConversationResponse mapping
type SummarizeConversationResponse struct {
  Summary  string  `json:"summary"`
}


// String returns the string representation of SummarizeParagraphsResponse
func (instance SummarizeParagraphsResponse) String() string {
  return Stringify(instance)
}

// String returns the string representation of SummarizeConversationResponse
func (instance SummarizeConversationResponse) String() string {
  return Stringify(instance)
}


// SummarizeParagraphs summarize given paragraphs.
func (service *TaskService) SummarizeParagraphs(data SummarizeParagraphsRequest) (*SummarizeParagraphsResponse, error) {
  req, _ := service.client.NewRequest("POST", "task/summarize/paragraphs", data)

  result := new(SummarizeParagraphsResponseData)
  _, err := service.client.Do(req, result)
  if err != nil {
    return nil, err
  }

  return result.Data, err
}


// SummarizeConversation summarize a given conversation, from a list of messages.
func (service *TaskService) SummarizeConversation(data SummarizeConversationRequest) (*SummarizeConversationResponse, error) {
  req, _ := service.client.NewRequest("POST", "task/summarize/conversation", data)

  result := new(SummarizeConversationResponseData)
  _, err := service.client.Do(req, result)
  if err != nil {
    return nil, err
  }

  return result.Data, err
}
