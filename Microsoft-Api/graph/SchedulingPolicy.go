package graph

import(
  "time"
)

type SchedulingPolicy struct{
  allowStaffSelection      bool      `json "allowStaffSelection"`
  MaximumAdvance           time.Time `json "maximumAdvance"`
  MinimumLeadTime          time.Time `json "minimumLeadTime"`
  SendConfirmationsToOwner bool      `json "sendConfirmationsToOwner"`
  TimeSlotInterval         time.Time `json "timeSlotInterval"`
}
