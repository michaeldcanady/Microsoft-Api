package graph

import(

)

type DriveItem struct{
  audio             Audio
  content": { "@odata.type": "Edm.Stream" },
  cTag              string
  deleted           Deleted
  description       string
  file              File
  fileSystemInfo    FileSystemInfo
  folder            Folder
  image             Image
  location          GeoCoordinates
  package           Package
  pendingOperations PendingOperations
  photo             Photo
  publication       PublicationFacet
  remoteItem": { "@odata.type": "microsoft.graph.remoteItem" },
  root": { "@odata.type": "microsoft.graph.root" },
  searchResult": { "@odata.type": "microsoft.graph.searchResult" },
  shared": { "@odata.type": "microsoft.graph.shared" },
  sharepointIds": { "@odata.type": "microsoft.graph.sharepointIds" },
  size": 1024,
  specialFolder": { "@odata.type": "microsoft.graph.specialFolder" },
  video": { "@odata.type": "microsoft.graph.video" },
  webDavUrl": "string",

  /* relationships */
  activities": [{"@odata.type": "microsoft.graph.itemActivity"}],
  analytics": {"@odata.type": "microsoft.graph.itemAnalytics"},
  children": [{ "@odata.type": "microsoft.graph.driveItem" }],
  createdByUser": { "@odata.type": "microsoft.graph.user" },
  lastModifiedByUser": { "@odata.type": "microsoft.graph.user" },
  permissions": [ {"@odata.type": "microsoft.graph.permission"} ],
  subscriptions": [ {"@odata.type": "microsoft.graph.subscription"} ],
  thumbnails": [ {"@odata.type": "microsoft.graph.thumbnailSet"}],
  versions": [ {"@odata.type": "microsoft.graph.driveItemVersion"}],

  /* inherited from baseItem */
  id": "string (identifier)",
  createdBy": {"@odata.type": "microsoft.graph.identitySet"},
  createdDateTime": "String (timestamp)",
  eTag": "string",
  lastModifiedBy": {"@odata.type": "microsoft.graph.identitySet"},
  lastModifiedDateTime": "String (timestamp)",
  name": "string",
  parentReference": {"@odata.type": "microsoft.graph.itemReference"},
  webUrl": "string",

  /* instance annotations */
  "@microsoft.graph.conflictBehavior": "string",
  "@microsoft.graph.downloadUrl": "url",
  "@microsoft.graph.sourceUrl": "url"
}
