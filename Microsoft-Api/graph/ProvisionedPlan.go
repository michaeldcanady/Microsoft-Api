package graph

import(

)

type ProvisionedPlan struct{
  CapabilityStatus   string `json "capabilityStatus"`
  ProvisioningStatus string `json "provisioningStatus"`
  Service            string `json "service"`
}
