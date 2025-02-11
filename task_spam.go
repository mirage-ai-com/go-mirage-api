// go-mirage-api
//
// Copyright 2024, Valerian Saliou
// Author: Valerian Saliou <valerian@valeriansaliou.name>

package mirage


// SpamClassifyRequest mapping
type SpamClassifyRequest struct {
  Sender      SpamClassifyRequestSender        `json:"sender"`
  Transcript  []SpamClassifyRequestTranscript  `json:"transcript"`
}

// SpamClassifyRequestSender mapping
type SpamClassifyRequestSender struct {
  Name   *string  `json:"name"`
  Email  *string  `json:"email"`
}

// SpamClassifyRequestTranscript mapping
type SpamClassifyRequestTranscript struct {
  From    string  `json:"from"`
  Origin  string  `json:"origin"`
  Text    string  `json:"text"`
}


// SpamClassifyResponseData mapping
type SpamClassifyResponseData struct {
  Data  *SpamClassifyResponse  `json:"data"`
}

// SpamClassifyResponse mapping
type SpamClassifyResponse struct {
  Class       string                      `json:"class"`
  Confidence  float32                     `json:"confidence"`
  LogProb     float32                     `json:"logprob"`
  Scores      SpamClassifyResponseScores  `json:"scores"`
}

// SpamClassifyResponseScores mapping
type SpamClassifyResponseScores struct {
  Gibberish  float32  `json:"gibberish"`
  Marketing  float32  `json:"marketing"`
  Regular    float32  `json:"regular"`
  Spam       float32  `json:"spam"`
}


// String returns the string representation of SpamClassifyResponse
func (instance SpamClassifyResponse) String() string {
  return Stringify(instance)
}


// SpamClassify spam check classification on spammy emails using a sender name, sender email and transcript.
func (service *TaskService) SpamClassify(ctx RequestContext, data SpamClassifyRequest) (*SpamClassifyResponse, error) {
  req, _ := service.client.NewRequest("POST", "task/spam/classify", data, ctx)

  result := new(SpamClassifyResponseData)
  _, err := service.client.Do(req, result)
  if err != nil {
    return nil, err
  }

  return result.Data, err
}
