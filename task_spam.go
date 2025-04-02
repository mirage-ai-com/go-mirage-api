// go-mirage-api
//
// Copyright 2024, Valerian Saliou
// Author: Valerian Saliou <valerian@valeriansaliou.name>

package mirage


// SpamConversationRequest mapping
type SpamConversationRequest struct {
  Sender      SpamConversationRequestSender        `json:"sender"`
  Transcript  []SpamConversationRequestTranscript  `json:"transcript"`
}

// SpamConversationRequestSender mapping
type SpamConversationRequestSender struct {
  Name   *string  `json:"name"`
  Email  *string  `json:"email"`
}

// SpamConversationRequestTranscript mapping
type SpamConversationRequestTranscript struct {
  From    string  `json:"from"`
  Origin  string  `json:"origin"`
  Text    string  `json:"text"`
}

// SpamDocumentRequest mapping
type SpamDocumentRequest struct {
  Name     string  `json:"name"`
  Domain   string  `json:"domain"`
  Title    string  `json:"title"`
  Content  string  `json:"content"`
}


// SpamGenericResponseData mapping
type SpamGenericResponseData struct {
  Data  *SpamGenericResponse  `json:"data"`
}

// SpamGenericResponse mapping
type SpamGenericResponse struct {
  Class       string                     `json:"class"`
  Confidence  float32                    `json:"confidence"`
  LogProb     float32                    `json:"logprob"`
  Scores      SpamGenericResponseScores  `json:"scores"`
}

// SpamGenericResponseScores mapping
type SpamGenericResponseScores struct {
  Gibberish  *float32  `json:"gibberish,omitempty"`
  Marketing  *float32  `json:"marketing,omitempty"`
  Regular    *float32  `json:"regular,omitempty"`
  Spam       float32   `json:"spam"`
}


// String returns the string representation of SpamGenericResponse
func (instance SpamGenericResponse) String() string {
  return Stringify(instance)
}


// SpamConversation spam check classification for conversations on spammy emails using a sender name, sender email and transcript.
func (service *TaskService) SpamConversation(ctx RequestContext, data SpamConversationRequest) (*SpamGenericResponse, error) {
  req, _ := service.client.NewRequest("POST", "task/spam/conversation", data, ctx)

  result := new(SpamGenericResponseData)
  _, err := service.client.Do(req, result)
  if err != nil {
    return nil, err
  }

  return result.Data, err
}


// SpamDocument spam check classification for documents on spammy documents using a title, content, author name and author domain.
func (service *TaskService) SpamDocument(ctx RequestContext, data SpamDocumentRequest) (*SpamGenericResponse, error) {
  req, _ := service.client.NewRequest("POST", "task/spam/document", data, ctx)

  result := new(SpamGenericResponseData)
  _, err := service.client.Do(req, result)
  if err != nil {
    return nil, err
  }

  return result.Data, err
}
