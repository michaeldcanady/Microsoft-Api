package graph

import(
  "time"
)

type Event struct{
  AllowNewTimeProposals         bool                                `json "allowNewTimeProposals"`
  Attendees                     []Attendee                          `json "attendees"`
  Body                          ItemBody                            `json "body"`
  BodyPreview                   string                              `json "bodyPreview"`
  Categories                    []string                            `json "categories"`
  ChangeKey                     string                              `json "changeKey"`
  CreatedDateTime               time.Time                           `json "createdDateTime"`
  End                           graph.dateTimeTimeZone              `json "end"`
  HasAttachments                bool                                `json "hasAttachments"`
  ICalUId                       string                              `json "iCalUId"`
  Id                            string                              `json "id"`
  Importance                    string                              `json "importance"`
  IsAllDay                      bool                                `json "isAllDay"`
  IsCancelled                   bool                                `json "isCancelled"`
  IsOnlineMeeting               bool                                `json "isOnlineMeeting"`
  IsOrganizer                   bool                                `json "isOrganizer"`
  IsReminderOn                  bool                                `json "isReminderOn"`
  LastModifiedDateTime          time.Time                           `json "lastModifiedDateTime"`
  Location                      Location                            `json "location"`
  Locations                     []Location                          `json "locations"`
  OnlineMeeting                 OnlineMeetingInfo                   `json "onlineMeeting"`
  OnlineMeetingProvider         string                              `json "onlineMeetingProvider"`
  OnlineMeetingUrl              string                              `json "onlineMeetingUrl"`
  Organizer                     Recipient                           `json "organizer"`
  OriginalEndTimeZone           string                              `json "originalEndTimeZone"`
  OriginalStart                 time.Time                           `json "originalStart"`
  OriginalStartTimeZone         string                              `json "originalStartTimeZone"`
  Recurrence                    PatternedRecurrence                 `json "recurrence"`
  ReminderMinutesBeforeStart    int                                 `json "reminderMinutesBeforeStart"`
  ResponseRequested             bool                                `json "responseRequested"`
  ResponseStatus                responseStatus                      `json "responseStatus"`
  Sensitivity                   string                              `json "sensitivity"`
  SeriesMasterId                string                              `json "seriesMasterId"`
  ShowAs                        string                              `json "showAs"`
  Start                         DateTimeTimeZone                    `json "start"`
  Subject                       string                              `json "subject"`
  TransactionId                 string                              `json "transactionId"`
  Type                          string                              `json "type"`
  WebLink                       string                              `json "webLink"`
  Attachments                   []Attachment                        `json "attachments"`
  Calendar                      Calendar                            `json "calendar"`
  Extensions                    []Extension                         `json "extensions"`
  Instances                     []Event                             `json "instances"`
  MultiValueExtendedProperties  []MultiValueLegacyExtendedProperty  `json "multiValueExtendedProperties"`
  SingleValueExtendedProperties []SingleValueLegacyExtendedProperty `json "singleValueExtendedProperties"`
}
