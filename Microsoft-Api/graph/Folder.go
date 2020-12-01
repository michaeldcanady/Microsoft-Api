package graph

import(

)

type Folder struct{
  ChildCount int        `json "childCount"`
  View       FolderView `json "view"`
}
