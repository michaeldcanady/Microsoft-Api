package graph

import(

)

type Identity struct{
  DisplayName string       `json "displayName"`
  Id          string       `json "id"`
  Thumbnails  ThumbnailSet `json "thumbnails"`
}
