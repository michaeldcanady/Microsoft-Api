package graph

import(
  "net/url"
  "time"
)

type Drive struct{
  id                   string
  createdBy            IdentitySet
  createdDateTime      time.Time
  description          string
  driveType": "personal | business | documentLibrary",
  following": [{"@odata.type": "microsoft.graph.driveItem"}],
  items": [ { "@odata.type": "microsoft.graph.driveItem" } ],
  lastModifiedBy       IdentitySet
  lastModifiedDateTime time.Time
  name                 string
  owner                IdentitySet
  quota": { "@odata.type": "microsoft.graph.quota" },
  root": { "@odata.type": "microsoft.graph.driveItem" },
  sharepointIds": { "@odata.type": "microsoft.graph.sharepointIds" },
  special": [ { "@odata.type": "microsoft.graph.driveItem" }],
  system": { "@odata.type": "microsoft.graph.systemFacet" },
  webUrl url.URL
}
