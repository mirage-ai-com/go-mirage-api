// go-mirage-api
//
// Copyright 2023, Valerian Saliou
// Author: Valerian Saliou <valerian@valeriansaliou.name>

package mirage

// TranscribeSpeechRequest mapping
type TranscribeSpeechRequest struct {
	Locale *TranscribeSpeechRequestLocale `json:"locale,omitempty"`
	Media  TranscribeSpeechRequestMedia   `json:"media"`
}

// TranscribeSpeechRequestLocale mapping
type TranscribeSpeechRequestLocale struct {
	To string `json:"to"`
}

// TranscribeSpeechRequestMedia mapping
type TranscribeSpeechRequestMedia struct {
	Type *string `json:"type,omitempty"`
	URL  string  `json:"url"`
}

// TranscribeSpeechResponseData mapping
type TranscribeSpeechResponseData struct {
	Data *TranscribeSpeechResponse `json:"data"`
}

// TranscribeSpeechResponse mapping
type TranscribeSpeechResponse struct {
	Locale string                         `json:"locale"`
	Parts  []TranscribeSpeechResponsePart `json:"parts"`
}

// TranscribeSpeechResponse mapping
type TranscribeSpeechResponsePart struct {
	Start float32 `json:"start"`
	End   float32 `json:"end"`
	Text  string  `json:"text"`
}

// String returns the string representation of TranscribeSpeechResponse
func (instance TranscribeSpeechResponse) String() string {
	return Stringify(instance)
}

// TranscribeSpeech transcribe speech from an audio file to text.
func (service *TaskService) TranscribeSpeech(ctx RequestContext, data TranscribeSpeechRequest) (*TranscribeSpeechResponse, error) {
	req, _ := service.client.NewRequest("POST", "task/transcribe/speech", data, ctx)

	result := new(TranscribeSpeechResponseData)
	_, err := service.client.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result.Data, err
}
