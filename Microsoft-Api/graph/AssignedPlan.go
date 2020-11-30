package graph

import(
  "time"
)

type AssignedPlan struct{
  AssignedDateTime time.Time `json "assignedDateTime"`
  CapalilityStatus string    `json "capalilityStatus"`
  Service          string    `json "service"`
  ServicePlanId    string    `json "servicePlanId"`
}
