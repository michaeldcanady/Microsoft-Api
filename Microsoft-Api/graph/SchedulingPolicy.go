package graph

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
