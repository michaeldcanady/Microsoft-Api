package graph

import(
  "time"
)

type AssignedLicense struct{
  DisabledPlans []string `json "disabledPlans"`
  SkuId         string   `json "skuId"`
}

type AssignedPlan struct{
  AssignedDateTime time.Time `json "assignedDateTime"`
  CapalilityStatus string    `json "capalilityStatus"`
  Service          string    `json "service"`
  ServicePlanId    string    `json "servicePlanId"`
}

type EmployeeOrgData struct{
  CostCenter string `json "costCenter"`
  Division   string `json "division"`
}

type ObjectIdentity struct{
  SignInType       string `json "signInType"`
  Issuer           string `json "issuer"`
  IssuerAssignedId string `json "issuerAssignedId"`
}

type LicenseAssignmentState struct{
  AssignedByGroup string `json "assignedByGroup"`
  DisabledPlans   string `json "disabledPlans"`
  Error           string `json "error"`
  SkuId           string `json "skuId"`
  State           string `json "state"`
}

type MailboxSettings struct{
  archiveFolder                         string                  `json "archiveFolder"`
  automaticRepliesSetting               automaticRepliesSetting `json "automaticRepliesSetting"`
  dateFormat                            string                  `json "dateFormat"`
  delegateMeetingMessageDeliveryOptions string                  `json "delegateMeetingMessageDeliveryOptions"`
  language                              localeInfo              `json "language"`
  timeFormat                            string                  `json "timeFormat"`
  timeZone                              string                  `json "timeZone"`
  workingHours                          WorkingHours            `json "workingHours"`
}

type automaticRepliesSetting struct{
  ExternalAudience       string           `json "externalAudience"`
  ExternalReplyMessage   string           `json "externalReplyMessage"`
  InternalReplyMessage   string           `json "internalReplyMessage"`
  ScheduledEndDateTime   DateTimeTimeZone `json "scheduledEndDateTime"`
  ScheduledStartDateTime DateTimeTimeZone `json "scheduledStartDateTime"`
  Status                 string           `json "status"`
}

type DateTimeTimeZone struct{
  DateTime string `json "dateTime"`
  TimeZone string `json "timeZone"`
}

type WorkingHours struct{
  DaysOfWeek []string     `json "daysOfWeek"`
  StartTime  time.Time    `json "startTime"`
  EndTime    time.Time    `json "endTime"`
  TimeZone   TimeZoneBase `json "timeZone"`
}

type TimeZoneBase struct{
  Name string `json "name"`
}

type localeInfo struct{
  Locale      string `json "locale"`
  DisplayName string `json "displayName"`
}

type OnPremisesExtensionAttributes struct{
  extensionAttribute1  string `json "extensionAttribute1"`
  extensionAttribute2  string `json "extensionAttribute2"`
  extensionAttribute3  string `json "extensionAttribute3"`
  extensionAttribute4  string `json "extensionAttribute4"`
  extensionAttribute5  string `json "extensionAttribute5"`
  extensionAttribute6  string `json "extensionAttribute6"`
  extensionAttribute7  string `json "extensionAttribute7"`
  extensionAttribute8  string `json "extensionAttribute8"`
  extensionAttribute9  string `json "extensionAttribute9"`
  extensionAttribute10 string `json "extensionAttribute10"`
  extensionAttribute11 string `json "extensionAttribute11"`
  extensionAttribute12 string `json "extensionAttribute12"`
  extensionAttribute13 string `json "extensionAttribute13"`
  extensionAttribute14 string `json "extensionAttribute14"`
  extensionAttribute15 string `json "extensionAttribute15"`
}

type onPremisesProvisioningError struct{
  category             string    `json "category"`
  occurredDateTime     time.Time `json "occurredDateTime"`
  propertyCausingError string    `json "propertyCausingError"`
  value                string    `json "value"`
}

type passwordProfile struct{
  forceChangePasswordNextSignIn        bool   `json "forceChangePasswordNextSignIn"`
  forceChangePasswordNextSignInWithMfa bool   `json "forceChangePasswordNextSignInWithMfa"`
  password                             string `json "password"`
}

type provisionedPlan struct{
  capabilityStatus   string `json "capabilityStatus"`
  provisioningStatus string `json "provisioningStatus"`
  service            string `json "service"`
}

type Calendar struct{
  allowedOnlineMeetingProviders []string     `json "allowedOnlineMeetingProviders"`
  canEdit                       bool         `json "canEdit"`
  canShare                      bool         `json "canShare"`
  canViewPrivateItems           bool         `json "canViewPrivateItems"`
  changeKey                     string       `json "changeKey"`
  color                         string       `json "color"`
  defaultOnlineMeetingProvider  string       `json "defaultOnlineMeetingProvider"`
  id                            string       `json "id"`
  isRemovable                   bool         `json "isRemovable"`
  isTallyingResponses           bool         `json "isTallyingResponses"`
  name                          string       `json "name"`
  owner                         EmailAddress `json "owner"`
}

type CalendarGroup struct{
  changeKey string `json "changeKey"`
  classId   string `json "classId"`
  id        string `json "id"`
  name      string `json "name"`
}
