// go-mirage-api
//
// Copyright 2023, Valerian Saliou
// Author: Valerian Saliou <valerian@valeriansaliou.name>

package mirage


// SummarizeParagraphsRequest mapping
type SummarizeParagraphsRequest struct {
  Locale      *SummarizeParagraphsRequestLocale      `json:"locale,omitempty"`
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
  Locale      *SummarizeConversationRequestLocale       `json:"locale,omitempty"`
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


// SummarizeGenericResponseData mapping
type SummarizeGenericResponseData struct {
  Data  *SummarizeGenericResponse  `json:"data"`
}

// SummarizeGenericResponse mapping
type SummarizeGenericResponse struct {
  Summary  string  `json:"summary"`
}


// String returns the string representation of SummarizeGenericResponse
func (instance SummarizeGenericResponse) String() string {
  return Stringify(instance)
}


// SummarizeParagraphs summarize given paragraphs.
func (service *TaskService) SummarizeParagraphs(ctx RequestContext, data SummarizeParagraphsRequest) (*SummarizeGenericResponse, error) {
  req, _ := service.client.NewRequest("POST", "task/summarize/paragraphs", data, ctx)

  result := new(SummarizeGenericResponseData)
  _, err := service.client.Do(req, result)
  if err != nil {
    return nil, err
  }

  return result.Data, err
}


// SummarizeConversation summarize a given conversation, from a list of messages.
func (service *TaskService) SummarizeConversation(ctx RequestContext, data SummarizeConversationRequest) (*SummarizeGenericResponse, error) {
  req, _ := service.client.NewRequest("POST", "task/summarize/conversation", data, ctx)

  result := new(SummarizeGenericResponseData)
  _, err := service.client.Do(req, result)
  if err != nil {
    return nil, err
  }

  return result.Data, err
}
