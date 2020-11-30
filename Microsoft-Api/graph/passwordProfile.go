package graph

import(

)

type PasswordProfile struct{
  ForceChangePasswordNextSignIn        bool   `json "forceChangePasswordNextSignIn"`
  ForceChangePasswordNextSignInWithMfa bool   `json "forceChangePasswordNextSignInWithMfa"`
  Password                             string `json "password"`
}
