package graph

import(

)

type FolderView struct{
  SortBy    string `json "sortBy"`
  SortOrder string `json "sortOrder"`
  ViewType  string `json "viewType"`
}
