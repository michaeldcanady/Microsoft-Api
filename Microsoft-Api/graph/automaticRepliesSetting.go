package graph

import(

)

type AutomaticRepliesSetting struct{
  ExternalAudience       string           `json "externalAudience"`
  ExternalReplyMessage   string           `json "externalReplyMessage"`
  InternalReplyMessage   string           `json "internalReplyMessage"`
  ScheduledEndDateTime   DateTimeTimeZone `json "scheduledEndDateTime"`
  ScheduledStartDateTime DateTimeTimeZone `json "scheduledStartDateTime"`
  Status                 string           `json "status"`
}
