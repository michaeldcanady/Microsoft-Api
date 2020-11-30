package main

import(
  "time"
  "github.com/michaeldcanady/GoOutlook/GoOutlook"
)

type SchedulingPolicy struct{
  allowStaffSelection bool
  maximumAdvance time.Time
  minimumLeadTime time.Time
  sendConfirmationsToOwner bool
  timeSlotInterval time.Time
}

type Location struct{
  Address PhysicalAddress
  Coordinates outlook.GeoCoordinates
  DisplayName string
  LocationEmailAddress string
  LocationUri string
  LocationType string
  UniqueId string
  UniqueIdType string
}

type WorkHours struct{
  Day string
  TimeSlots []WorkTimeSlot
}

type WorkTimeSlot struct{
  Start time.Time
  End time.Time
}

type Reminder struct{
  Message string
  Offset time.Time
  Recipients string
}

type PhysicalAddress struct{
  City string
  CountryOrRegion string
  PostalCode string
  State string
  Street string
}
