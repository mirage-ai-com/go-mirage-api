// go-mirage-api
//
// Copyright 2023, Valerian Saliou
// Author: Valerian Saliou <valerian@valeriansaliou.name>

package mirage


// AnswerPromptRequest mapping
type AnswerPromptRequest struct {
  Prompt  string                      `json:"prompt"`
  Answer  *AnswerPromptRequestAnswer  `json:"answer,omitempty"`
  Schema  *interface{}                `json:"schema,omitempty"`
  Model   *string                     `json:"model,omitempty"`
}

// AnswerPromptRequestAnswer mapping
type AnswerPromptRequestAnswer struct {
  MaxTokens    *uint16   `json:"max_tokens,omitempty"`
  Temperature  *float32  `json:"temperature,omitempty"`
}

// AnswerQuestionRequest mapping
type AnswerQuestionRequest struct {
  Question  string                        `json:"question"`
  Answer    *AnswerQuestionRequestAnswer  `json:"answer,omitempty"`
  Locale    *AnswerQuestionRequestLocale  `json:"locale,omitempty"`
  Context   AnswerQuestionRequestContext  `json:"context"`
  Model     *string                       `json:"model,omitempty"`
}

// AnswerQuestionRequestAnswer mapping
type AnswerQuestionRequestAnswer struct {
  Start        *string   `json:"start,omitempty"`
  System       *string   `json:"system,omitempty"`
  Quality      *uint8    `json:"quality,omitempty"`
  MaxTokens    *uint16   `json:"max_tokens,omitempty"`
  Temperature  *float32  `json:"temperature,omitempty"`
}

// AnswerQuestionRequestLocale mapping
type AnswerQuestionRequestLocale struct {
  From  string  `json:"from"`
}

// AnswerQuestionRequestContext mapping
type AnswerQuestionRequestContext struct {
  Source        *string                                   `json:"source,omitempty"`
  PrimaryID     string                                    `json:"primary_id"`
  Filters       *AnswerQuestionRequestContextFilters      `json:"filters,omitempty"`
  Conversation  AnswerQuestionRequestContextConversation  `json:"conversation"`
}

// AnswerQuestionRequestContextFilters mapping
type AnswerQuestionRequestContextFilters struct {
  SecondaryID  *AnswerQuestionRequestContextFiltersFilter  `json:"secondary_id,omitempty"`
  TertiaryID   *AnswerQuestionRequestContextFiltersFilter  `json:"tertiary_id,omitempty"`
  Source       *AnswerQuestionRequestContextFiltersFilter  `json:"source,omitempty"`
}

// AnswerQuestionRequestContextFiltersFilter mapping
type AnswerQuestionRequestContextFiltersFilter struct {
  Include  *[]string  `json:"include,omitempty"`
  Exclude  *[]string  `json:"exclude,omitempty"`
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

// AnswerChatRequest mapping
type AnswerChatRequest struct {
  Schema   *interface{}              `json:"schema,omitempty"`
  Model    *string                   `json:"model,omitempty"`
  Context  AnswerChatRequestContext  `json:"context"`
  Tools    []AnswerChatRequestTool   `json:"tools,omitempty"`
}

// AnswerChatRequestContext mapping
type AnswerChatRequestContext struct {
  Conversation  AnswerChatRequestContextConversation  `json:"conversation"`
}

// AnswerChatRequestContextConversation mapping
type AnswerChatRequestContextConversation struct {
	Messages  []AnswerChatRequestContextConversationMessage  `json:"messages"`
}

// AnswerChatRequestContextConversationMessage mapping
type AnswerChatRequestContextConversationMessage struct {
	From  string  `json:"from"`
	Text  string  `json:"text"`
}

// AnswerChatRequestTool mapping
type AnswerChatRequestTool struct {
	Type      string  `json:"type"`
	Function  AnswerChatRequestToolFunction  `json:"function"`
}

// AnswerChatRequestToolFunction mapping
type AnswerChatRequestToolFunction struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Parameters  *interface{}  `json:"parameters,omitempty"`
}

// AnswerPromptResponseData mapping
type AnswerPromptResponseData struct {
  Data  *AnswerPromptResponse  `json:"data"`
}

// AnswerPromptResponse mapping
type AnswerPromptResponse struct {
  Answer  string  `json:"answer"`
  Model   string  `json:"model"`
}

// AnswerQuestionResponseData mapping
type AnswerQuestionResponseData struct {
  Data  *AnswerQuestionResponse  `json:"data"`
}

// AnswerQuestionResponse mapping
type AnswerQuestionResponse struct {
  Answer   string                          `json:"answer"`
  Model    string                          `json:"model"`
  Sources  []AnswerQuestionResponseSource  `json:"sources"`
}

// AnswerQuestionResponseSource mapping
type AnswerQuestionResponseSource struct {
  Source       *string             `json:"source,omitempty"`
  Score        *uint8              `json:"score,omitempty"`
  PrimaryID    string              `json:"primary_id"`
  SecondaryID  *string             `json:"secondary_id,omitempty"`
  Excerpt      *string             `json:"excerpt,omitempty"`
  Timestamp    *uint64             `json:"timestamp,omitempty"`
  Metadata     *map[string]string  `json:"metadata,omitempty"`
}

// AnswerPromptResponseData mapping
type AnswerChatResponseData struct {
  Data  *AnswerChatResponse  `json:"data"`
}

// AnswerChatResponse mapping
type AnswerChatResponse struct {
  Answer     string                         `json:"answer"`
  Model      string                         `json:"model"`
  ToolCalls  *[]AnswerChatResponseToolCall  `json:"tool_calls,omitempty"`
}

// AnswerChatResponseToolCall mapping
type AnswerChatResponseToolCall struct {
  Function  *AnswerChatResponseToolCallFunction  `json:"function,omitempty"`
}

// AnswerChatResponseToolCallFunction mapping
type AnswerChatResponseToolCallFunction struct {
  Name       string        `json:"name"`
  Arguments  *interface{}  `json:"arguments,omitempty"`
}

// String returns the string representation of AnswerPromptResponse
func (instance AnswerPromptResponse) String() string {
  return Stringify(instance)
}

// String returns the string representation of AnswerQuestionResponse
func (instance AnswerQuestionResponse) String() string {
  return Stringify(instance)
}

// String returns the string representation of AnswerChatResponse
func (instance AnswerChatResponse) String() string {
  return Stringify(instance)
}

// AnswerPrompt answer a given prompt.
func (service *TaskService) AnswerPrompt(ctx RequestContext, data AnswerPromptRequest) (*AnswerPromptResponse, error) {
  req, _ := service.client.NewRequest("POST", "task/answer/prompt", data, ctx)

  result := new(AnswerPromptResponseData)
  _, err := service.client.Do(req, result)
  if err != nil {
    return nil, err
  }

  return result.Data, err
}


// AnswerQuestion answer a given question.
func (service *TaskService) AnswerQuestion(ctx RequestContext, data AnswerQuestionRequest) (*AnswerQuestionResponse, error) {
  req, _ := service.client.NewRequest("POST", "task/answer/question", data, ctx)

  result := new(AnswerQuestionResponseData)
  _, err := service.client.Do(req, result)
  if err != nil {
    return nil, err
  }

  return result.Data, err
}

// AnswerChat answer a given chat.
func (service *TaskService) AnswerChat(ctx RequestContext, data AnswerChatRequest) (*AnswerChatResponse, error) {
  req, _ := service.client.NewRequest("POST", "task/answer/chat", data, ctx)

  result := new(AnswerChatResponseData)
  _, err := service.client.Do(req, result)
  if err != nil {
    return nil, err
  }

  return result.Data, err
}
