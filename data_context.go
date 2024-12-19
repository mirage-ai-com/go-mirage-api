// go-mirage-api
//
// Copyright 2023, Valerian Saliou
// Author: Valerian Saliou <valerian@valeriansaliou.name>

package mirage


// IngestContextDataRequest mapping
type IngestContextDataRequest struct {
  Items  []IngestContextDataRequestItem  `json:"items"`
}

// IngestContextDataRequestItem mapping
type IngestContextDataRequestItem struct {
  Operation     string              `json:"operation"`
  PrimaryID     string              `json:"primary_id"`
  SecondaryID   *string             `json:"secondary_id,omitempty"`
  TertiaryID    *string             `json:"tertiary_id,omitempty"`
  Text          *string             `json:"text,omitempty"`
  Timestamp     *uint64             `json:"timestamp,omitempty"`
  Source        *string             `json:"source,omitempty"`
  Metadata      *map[string]string  `json:"metadata,omitempty"`
}


// IngestContextDataResponseData mapping
type IngestContextDataResponseData struct {
  Data  *IngestContextDataResponse  `json:"data"`
}

// IngestContextDataResponse mapping
type IngestContextDataResponse struct {
  Imported  bool  `json:"imported"`
}


// String returns the string representation of IngestContextDataResponse
func (instance IngestContextDataResponse) String() string {
  return Stringify(instance)
}


// IngestContextData ingest context data into account.
func (service *TaskService) IngestContextData(ctx RequestContext, data IngestContextDataRequest) (*IngestContextDataResponse, error) {
  req, _ := service.client.NewRequest("POST", "data/context/ingest", data, ctx)

  result := new(IngestContextDataResponseData)
  _, err := service.client.Do(req, result)
  if err != nil {
    return nil, err
  }

  return result.Data, err
}
