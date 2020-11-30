package graph

import(

)

type Calendar struct{
  AllowedOnlineMeetingProviders []string     `json "allowedOnlineMeetingProviders"`
  CanEdit                       bool         `json "canEdit"`
  CanShare                      bool         `json "canShare"`
  CanViewPrivateItems           bool         `json "canViewPrivateItems"`
  ChangeKey                     string       `json "changeKey"`
  Color                         string       `json "color"`
  DefaultOnlineMeetingProvider  string       `json "defaultOnlineMeetingProvider"`
  Id                            string       `json "id"`
  IsRemovable                   bool         `json "isRemovable"`
  IsTallyingResponses           bool         `json "isTallyingResponses"`
  Name                          string       `json "name"`
  Owner                         EmailAddress `json "owner"`
}
