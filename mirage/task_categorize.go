// go-mirage-api
//
// Copyright 2023, Valerian Saliou
// Author: Valerian Saliou <valerian@valeriansaliou.name>

package mirage


// CategorizeConversationsRequest mapping
type CategorizeConversationsRequest struct {
  Conversations  []CategorizeConversationsRequestConversation  `json:"conversations"`
}

// CategorizeConversationsRequestConversation mapping
type CategorizeConversationsRequestConversation struct {
  Transcript  []CategorizeConversationsRequestConversationTranscript  `json:"transcript"`
}

// CategorizeConversationsRequestConversationTranscript mapping
type CategorizeConversationsRequestConversationTranscript struct {
  From  string  `json:"from"`
  Text  string  `json:"text"`
}


// CategorizeConversationsResponseData mapping
type CategorizeConversationsResponseData struct {
  Data  *CategorizeConversationsResponse  `json:"data"`
}

// CategorizeConversationsResponse mapping
type CategorizeConversationsResponse struct {
  Categories  []string  `json:"categories"`
}


// String returns the string representation of CategorizeConversationsResponse
func (instance CategorizeConversationsResponse) String() string {
  return Stringify(instance)
}


// CategorizeConversations categorize multiple conversations, from a list of messages for each individual conversation.
func (service *TaskService) CategorizeConversations(data CategorizeConversationsRequest) (*CategorizeConversationsResponse, error) {
  req, _ := service.client.NewRequest("POST", "task/categorize/conversations", data)

  result := new(CategorizeConversationsResponseData)
  _, err := service.client.Do(req, result)
  if err != nil {
    return nil, err
  }

  return result.Data, err
}
