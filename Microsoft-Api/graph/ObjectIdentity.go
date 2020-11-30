package graph

import(

)

type ObjectIdentity struct{
  SignInType       string `json "signInType"`
  Issuer           string `json "issuer"`
  IssuerAssignedId string `json "issuerAssignedId"`
}
