package graph

import(

)

type Attendee struct{
  EmailAddress    EmailAddress   `json "emailAddress"`
  ProposedNewTime TimeSlot       `json "proposedNewTime"`
  Status          ResponseStatus `json "status"`
  Type            string         `json "type"`
}
