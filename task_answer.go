// go-mirage-api
//
// Copyright 2023, Valerian Saliou
// Author: Valerian Saliou <valerian@valeriansaliou.name>

package mirage


// AnswerQuestionRequest mapping
type AnswerQuestionRequest struct {
  Question  string  											 `json:"question"`
  Answer    *AnswerQuestionRequestAnswer   `json:"answer,omitempty"`
  Locale    *AnswerQuestionRequestLocale   `json:"locale,omitempty"`
  Context   AnswerQuestionRequestContext   `json:"context"`
}

// AnswerQuestionRequestAnswer mapping
type AnswerQuestionRequestAnswer struct {
  Start  *string  `json:"start,omitempty"`
}

// AnswerQuestionRequestLocale mapping
type AnswerQuestionRequestLocale struct {
  From  string  `json:"from"`
}

// AnswerQuestionRequestContext mapping
type AnswerQuestionRequestContext struct {
  Team         AnswerQuestionRequestContextTeam         `json:"team"`
  Transcripts  AnswerQuestionRequestContextTranscripts  `json:"transcripts"`
}

// AnswerQuestionRequestContextTeam mapping
type AnswerQuestionRequestContextTeam struct {
  ID    string  `json:"id"`
  Name  string  `json:"name"`
}

// AnswerQuestionRequestContextTranscripts mapping
type AnswerQuestionRequestContextTranscripts struct {
	Conversation  AnswerQuestionRequestContextTranscriptsConversation  `json:"conversation"`
	Related   		*[]AnswerQuestionRequestContextTranscriptsRelated    `json:"related,omitempty"`
}

// AnswerQuestionRequestContextTranscriptsConversation mapping
type AnswerQuestionRequestContextTranscriptsConversation struct {
	Messages  []AnswerQuestionRequestContextTranscript  `json:"messages"`
}

// AnswerQuestionRequestContextTranscriptsRelated mapping
type AnswerQuestionRequestContextTranscriptsRelated struct {
	Messages  []AnswerQuestionRequestContextTranscript  `json:"messages"`
}

// AnswerQuestionRequestContextTranscript mapping
type AnswerQuestionRequestContextTranscript struct {
	From  string  `json:"from"`
	Text  string  `json:"text"`
}


// AnswerQuestionResponseData mapping
type AnswerQuestionResponseData struct {
  Data  *AnswerQuestionResponse  `json:"data"`
}

// AnswerQuestionResponse mapping
type AnswerQuestionResponse struct {
  Answer  string  `json:"answer"`
}


// String returns the string representation of AnswerQuestionResponse
func (instance AnswerQuestionResponse) String() string {
  return Stringify(instance)
}


// AnswerQuestion answer a given question.
func (service *TaskService) AnswerQuestion(data AnswerQuestionRequest) (*AnswerQuestionResponse, error) {
  req, _ := service.client.NewRequest("POST", "task/answer/question", data)

  result := new(AnswerQuestionResponseData)
  _, err := service.client.Do(req, result)
  if err != nil {
    return nil, err
  }

  return result.Data, err
}
