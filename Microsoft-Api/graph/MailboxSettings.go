package graph

import(

)

type MailboxSettings struct{
  ArchiveFolder                         string                  `json "archiveFolder"`
  AutomaticRepliesSetting               automaticRepliesSetting `json "automaticRepliesSetting"`
  DateFormat                            string                  `json "dateFormat"`
  DelegateMeetingMessageDeliveryOptions string                  `json "delegateMeetingMessageDeliveryOptions"`
  Language                              localeInfo              `json "language"`
  TimeFormat                            string                  `json "timeFormat"`
  TimeZone                              string                  `json "timeZone"`
  WorkingHours                          WorkingHours            `json "workingHours"`
}
