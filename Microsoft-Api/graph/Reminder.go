package graph

import(
  "time"
)

type Reminder struct{
  Message    string    `json "message"`
  Offset     time.Time `json "offset"`
  Recipients string    `json "recipients"`
}
