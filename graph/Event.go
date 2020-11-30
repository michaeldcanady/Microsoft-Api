package graph

import(

)

type event struct{
  allowNewTimeProposals bool
  attendees             []Attendee
  body                  ItemBody
  bodyPreview           string
  categories            []string
  changeKey             string
  createdDateTime       time.Time
  end                   graph.dateTimeTimeZone
  hasAttachments        bool
  iCalUId               string
  id                    string
  importance            string
  isAllDay              bool
  isCancelled           bool
  isOnlineMeeting       bool
  isOrganizer           bool
  isReminderOn          bool
  lastModifiedDateTime  time.Time
  location              Location
  locations             []Location
  onlineMeeting         OnlineMeetingInfo
  onlineMeetingProvider string
  onlineMeetingUrl      string
  organizer             Recipient
  originalEndTimeZone   string
  originalStart         time.Time
  originalStartTimeZone string
  recurrence": {"@odata.type": "microsoft.graph.patternedRecurrence"},
  reminderMinutesBeforeStart": 1024,
  responseRequested": true,
  responseStatus": {"@odata.type": "microsoft.graph.responseStatus"},
  sensitivity": "String",
  seriesMasterId": "string",
  showAs": "String",
  start": {"@odata.type": "microsoft.graph.dateTimeTimeZone"},
  subject": "string",
  transactionId": "string",
  type": "String",
  webLink": "string",
  attachments": [ { "@odata.type": "microsoft.graph.attachment" } ],
  calendar": { "@odata.type": "microsoft.graph.calendar" },
  extensions": [ { "@odata.type": "microsoft.graph.extension" } ],
  instances": [ { "@odata.type": "microsoft.graph.event" }],
  multiValueExtendedProperties": [ { "@odata.type": "microsoft.graph.multiValueLegacyExtendedProperty" }],
  singleValueExtendedProperties": [ { "@odata.type": "microsoft.graph.singleValueLegacyExtendedProperty" }]

}
