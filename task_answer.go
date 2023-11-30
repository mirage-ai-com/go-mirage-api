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
  Source        *string                                   `json:"source,omitempty"`
  PrimaryID     string                                    `json:"primary_id"`
  Conversation  AnswerQuestionRequestContextConversation  `json:"conversation"`
}

// AnswerQuestionRequestContextConversation mapping
type AnswerQuestionRequestContextConversation struct {
	Messages  []AnswerQuestionRequestContextConversationMessage  `json:"messages"`
}

// AnswerQuestionRequestContextConversationMessage mapping
type AnswerQuestionRequestContextConversationMessage struct {
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
