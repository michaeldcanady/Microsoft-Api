package graph

import(
  "time"
)

type RemoteItem struct{
  id                   string
  createdBy            IdentitySet
  createdDateTime      time.Time
  file                 File
  fileSystemInfo       FileSystemInfo
  folder               Folder
  image                Image
  lastModifiedBy       IdentitySet
  lastModifiedDateTime time.Time
  name                 string
  package              Package
  parentReference      ItemReference
  shared               Shared
  "sharepointIds": { "@odata.type": "microsoft.graph.sharepointIds" },
  "specialFolder": { "@odata.type": "microsoft.graph.specialFolder" },
  "size": 1024,
  "video": { "@odata.type": "microsoft.graph.video" },
  "webDavUrl": "url",
  "webUrl": "url"
}
