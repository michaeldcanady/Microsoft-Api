package graph

import(
  "time"
)

type AssignedLicense struct{
  DisabledPlans []string `json "disabledPlans"`
  SkuId string `json "skuId"`
}

type AssignedPlan struct{
  AssignedDateTime time.Time
  CapalilityStatus string
  Service string
  ServicePlanId String
}

type EmployeeOrgData struct{
  CostCenter string
  Division string
}

type ObjectIdentity struct{
  SignInType string
  Issuer string
  IssuerAssignedId string
}

type LicenseAssignmentState struct{
  AssignedByGroup string
  DisabledPlans string
  Error string
  SkuId string
  State string
}

type MailboxSettings struct{
  "archiveFolder": "string",
  "automaticRepliesSetting": {"@odata.type": "microsoft.graph.automaticRepliesSetting"},
  "dateFormat": "string",
  "delegateMeetingMessageDeliveryOptions": "String",
  "language": {"@odata.type": "microsoft.graph.localeInfo"},
  "timeFormat": "string",
  "timeZone": "string",
  "workingHours": {"@odata.type": "microsoft.graph.workingHours"}
}

