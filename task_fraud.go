// go-mirage-api
//
// Copyright 2023, Valerian Saliou
// Author: Valerian Saliou <valerian@valeriansaliou.name>

package mirage

// FraudSpamicityRequest mapping
type FraudSpamicityRequest struct {
	Name        string `json:"name"`
	Domain      string `json:"domain"`
	EmailDomain string `json:"email_domain"`
}

// FraudSpamicityResponseData mapping
type FraudSpamicityResponseData struct {
	Data *FraudSpamicityResponse `json:"data"`
}

// FraudSpamicityResponse mapping
type FraudSpamicityResponse struct {
	Fraud bool    `json:"fraud"`
	Score float32 `json:"score"`
}

// String returns the string representation of FraudSpamicityResponse
func (instance FraudSpamicityResponse) String() string {
	return Stringify(instance)
}

// FraudSpamicity fraud check classification on scammy websites using a website name, domain and email domain.
func (service *TaskService) FraudSpamicity(ctx RequestContext, data FraudSpamicityRequest) (*FraudSpamicityResponse, error) {
	req, _ := service.client.NewRequest("POST", "task/fraud/spamicity", data, ctx)

	result := new(FraudSpamicityResponseData)
	_, err := service.client.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result.Data, err
}
