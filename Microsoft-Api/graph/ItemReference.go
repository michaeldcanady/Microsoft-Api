package graph

import(
  "time"
)

type ItemReference struct{
  DriveId       string        `json "driveId"`
  DriveType     string        `json "driveType"`
  Id            string        `json "id"`
  Name          string        `json "name"`
  Path          string        `json "path"`
  ShareId       string        `json "shareId"`
  SharepointIds SharepointIds `json "sharepointIds"`
}
