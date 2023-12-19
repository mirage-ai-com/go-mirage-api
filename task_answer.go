// go-mirage-api
//
// Copyright 2023, Valerian Saliou
// Author: Valerian Saliou <valerian@valeriansaliou.name>

package mirage


// AnswerPromptRequest mapping
type AnswerPromptRequest struct {
  Prompt  string  `json:"prompt"`
}

// AnswerQuestionRequest mapping
type AnswerQuestionRequest struct {
  Question  string                         `json:"question"`
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

// AnswerPromptResponseData mapping
type AnswerPromptResponseData struct {
  Data  *AnswerPromptResponse  `json:"data"`
}

// AnswerPromptResponse mapping
type AnswerPromptResponse struct {
  Answer  string  `json:"answer"`
}

// AnswerQuestionResponseData mapping
type AnswerQuestionResponseData struct {
  Data  *AnswerQuestionResponse  `json:"data"`
}

// AnswerQuestionResponse mapping
type AnswerQuestionResponse struct {
  Answer  string  `json:"answer"`
}


// String returns the string representation of AnswerPromptResponse
func (instance AnswerPromptResponse) String() string {
  return Stringify(instance)
}

// String returns the string representation of AnswerQuestionResponse
func (instance AnswerQuestionResponse) String() string {
  return Stringify(instance)
}


// AnswerPrompt answer a given prompt.
func (service *TaskService) AnswerPrompt(data AnswerPromptRequest) (*AnswerPromptResponse, error) {
  req, _ := service.client.NewRequest("POST", "task/answer/prompt", data)

  result := new(AnswerPromptResponseData)
  _, err := service.client.Do(req, result)
  if err != nil {
    return nil, err
  }

  return result.Data, err
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
