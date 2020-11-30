package graph

import(

)

type IdentitySet struct{
  Application Identity `json "application"`
  Device      Identity `json "device"`
  User        Identity `json "user"`
}
