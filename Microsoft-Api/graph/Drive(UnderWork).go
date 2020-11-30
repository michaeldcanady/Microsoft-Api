package graph

import(
  "net/url"
  "time"
)

type Drive struct{
  id                   string
  createdBy": { "@odata.type": "microsoft.graph.identitySet" },
  createdDateTime      time.Time
  description          string
  driveType": "personal | business | documentLibrary",
  following": [{"@odata.type": "microsoft.graph.driveItem"}],
  items": [ { "@odata.type": "microsoft.graph.driveItem" } ],
  lastModifiedBy": { "@odata.type": "microsoft.graph.identitySet" },
  lastModifiedDateTime time.Time
  name                 string
  owner": { "@odata.type": "microsoft.graph.identitySet" },
  quota": { "@odata.type": "microsoft.graph.quota" },
  root": { "@odata.type": "microsoft.graph.driveItem" },
  sharepointIds": { "@odata.type": "microsoft.graph.sharepointIds" },
  special": [ { "@odata.type": "microsoft.graph.driveItem" }],
  system": { "@odata.type": "microsoft.graph.systemFacet" },
  webUrl url.URL
}
