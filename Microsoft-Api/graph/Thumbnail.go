package graph

import(

)

type Thumbnail struct{
  Height       int    `json "height"`
  SourceItemId string `json "sourceItemId"`
  Url          string `json "url"`
  Width        int    `json "width"`
  Content      []byte `json "content"`
}
