package graph

import(
  "time"
)

type responseStatus struct{
  Response String    `json "response"`
  Time     time.Time `json "time"`
}
