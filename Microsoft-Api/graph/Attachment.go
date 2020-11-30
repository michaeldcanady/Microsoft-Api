package graph

import(
  "time"
)

type Attachment struct{
  ContentType          string    `json "contentType"`
  Id                   string    `json "id"`
  IsInline             bool      `json "isInline"`
  LastModifiedDateTime time.Time `json "lastModifiedDateTime"`
  Name                 string    `json "name"`
  Size                 int       `json "size"`
}
