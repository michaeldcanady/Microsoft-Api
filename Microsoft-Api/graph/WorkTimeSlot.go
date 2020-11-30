package graph

import(
  "time"
)

type WorkTimeSlot struct{
  Start time.Time `json "start"`
  End   time.Time `json "end"`
}
