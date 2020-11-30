package graph

import(

)

type ThumbnailSet struct{
  Id     string    `json "id"`
  Large  Thumbnail `json "large"`
  Medium Thumbnail `json "medium"`
  Small  Thumbnail `json "small"`
  Source Thumbnail `json "source"`
}
