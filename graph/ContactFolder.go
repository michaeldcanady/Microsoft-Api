package graph

import(

)

type ContactFolder struct{
  DisplayName    string `json "displayName"`
  Id             string `json "id"`
  ParentFolderId string `json "parentFolderId"`
}
