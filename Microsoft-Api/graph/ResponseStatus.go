package graph

import(
  "time"
)

type responseStatus struct{
  response String    `json "response"`
  time     time.Time `json "time"`
}
