package graph

import(
  "time"
)

type PendingContentUpdate struct{
  QueuedDateTime time.Time `json "queuedDateTime"`
}
