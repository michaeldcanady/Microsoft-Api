package graph

import(

)

type Hashes struct{
  Crc32Hash    string `json "crc32Hash"`
  Sha1Hash     string `json "sha1Hash"`
  Sha256Hash   string `json "sha256Hash"`
  QuickXorHash string `json "quickXorHash"`
}
