package graph

import(

)

type File struct{
  Hashes   Hashes `json "hashes"`
  MimeType string `json "mimeType"`
}
