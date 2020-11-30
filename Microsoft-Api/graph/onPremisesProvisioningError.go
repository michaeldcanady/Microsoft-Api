package graph

import(
  "time"
)

type onPremisesProvisioningError struct{
  Category             string    `json "category"`
  OccurredDateTime     time.Time `json "occurredDateTime"`
  PropertyCausingError string    `json "propertyCausingError"`
  Value                string    `json "value"`
}
