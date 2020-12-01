package graph

import(
  "time"
)

type FileSystemInfo struct{
  CreatedDateTime      time.Time `json "createdDateTime"`
  LastAccessedDateTime time.Time `json "lastAccessedDateTime"`
  LastModifiedDateTime time.Time `json "lastModifiedDateTime"`
}
