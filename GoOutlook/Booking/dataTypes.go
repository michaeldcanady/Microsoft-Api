package main

import(
  "time"
)

type SchedulingPolicy struct{
  allowStaffSelection bool
  maximumAdvance time.Time
  minimumLeadTime time.Time
  sendConfirmationsToOwner bool
  timeSlotInterval time.Time
}

type Location struct{
  "address": {"@odata.type": "microsoft.graph.physicalAddress"},
  "coordinates": {"@odata.type": "microsoft.graph.outlookGeoCoordinates"},
  "displayName": "string",
  "locationEmailAddress": "string",
  "locationUri": "string",
  "locationType": "string",
  "uniqueId": "string",
  "uniqueIdType": "string"
}


type WorkHours struct{
  Day string
  TimeSlots []WorkTimeSlot
}

type WorkTimeSlot struct{
  Start time.Time
  End time.Time
}

type physicalAddress struct{
  City string
  CountryOrRegion string
  PostalCode string
  State string
  Street string
}
