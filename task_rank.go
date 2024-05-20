// go-mirage-api
//
// Copyright 2023, Valerian Saliou
// Author: Valerian Saliou <valerian@valeriansaliou.name>

package mirage


// RankQuestionRequest mapping
type RankQuestionRequest struct {
  Question  string                      `json:"question"`
  Context   RankQuestionRequestContext  `json:"context"`
}

// RankQuestionRequestContext mapping
type RankQuestionRequestContext struct {
  Source     *string                             `json:"source,omitempty"`
  PrimaryID  string                              `json:"primary_id"`
  Filters    *RankQuestionRequestContextFilters  `json:"filters,omitempty"`
}

// RankQuestionRequestContextFilters mapping
type RankQuestionRequestContextFilters struct {
  SecondaryID  *RankQuestionRequestContextFiltersFilter  `json:"secondary_id,omitempty"`
  TertiaryID   *RankQuestionRequestContextFiltersFilter  `json:"tertiary_id,omitempty"`
  Source       *RankQuestionRequestContextFiltersFilter  `json:"source,omitempty"`
}

// RankQuestionRequestContextFiltersFilter mapping
type RankQuestionRequestContextFiltersFilter struct {
  Include  *[]string  `json:"include,omitempty"`
  Exclude  *[]string  `json:"exclude,omitempty"`
}


// RankQuestionResponseData mapping
type RankQuestionResponseData struct {
  Data  *RankQuestionResponse  `json:"data"`
}

// RankQuestionResponse mapping
type RankQuestionResponse struct {
  Results  []RankQuestionResponseResults  `json:"results"`
}

// RankQuestionResponseResults mapping
type RankQuestionResponseResults struct {
  ID           string                             `json:"id"`
  Score        uint8                              `json:"score"`
  GroupedText  string                             `json:"grouped_text"`
  Items        []RankQuestionResponseResultsItem  `json:"items"`
}

// RankQuestionResponseResultsItem mapping
type RankQuestionResponseResultsItem struct {
  Source       *string       `json:"source,omitempty"`
  PrimaryID    *string       `json:"primary_id,omitempty"`
  SecondaryID  *string       `json:"secondary_id,omitempty"`
  Text         *string       `json:"text,omitempty"`
  Timestamp    *uint64       `json:"timestamp,omitempty"`
  Metadata     *interface{}  `json:"metadata,omitempty"`
}


// String returns the string representation of RankQuestionResponse
func (instance RankQuestionResponse) String() string {
  return Stringify(instance)
}


// RankQuestion ranks results based on a given question.
func (service *TaskService) RankQuestion(data RankQuestionRequest) (*RankQuestionResponse, error) {
  req, _ := service.client.NewRequest("POST", "task/rank/question", data)

  result := new(RankQuestionResponseData)
  _, err := service.client.Do(req, result)
  if err != nil {
    return nil, err
  }

  return result.Data, err
}
