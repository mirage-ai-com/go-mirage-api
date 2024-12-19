// go-mirage-api
//
// Copyright 2023, Valerian Saliou
// Author: Valerian Saliou <valerian@valeriansaliou.name>

package mirage


// TranslateTextRequest mapping
type TranslateTextRequest struct {
  Locale  TranslateTextRequestLocale  `json:"locale"`
  Type    *string                     `json:"type,omitempty"`
  Text    string                      `json:"text"`
}

// TranslateTextRequestLocale mapping
type TranslateTextRequestLocale struct {
  From  string  `json:"from"`
  To    string  `json:"to"`
}


// TranslateTextResponseData mapping
type TranslateTextResponseData struct {
  Data  *TranslateTextResponse  `json:"data"`
}

// TranslateTextResponse mapping
type TranslateTextResponse struct {
  Translation  string  `json:"translation"`
}


// String returns the string representation of TranslateTextResponse
func (instance TranslateTextResponse) String() string {
  return Stringify(instance)
}


// TranslateText translate a provided text between two languages.
func (service *TaskService) TranslateText(ctx RequestContext, data TranslateTextRequest) (*TranslateTextResponse, error) {
  req, _ := service.client.NewRequest("POST", "task/translate/text", data, ctx)

  result := new(TranslateTextResponseData)
  _, err := service.client.Do(req, result)
  if err != nil {
    return nil, err
  }

  return result.Data, err
}
