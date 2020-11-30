package graph

import(

)

type Calendar struct{
  allowedOnlineMeetingProviders []string     `json "allowedOnlineMeetingProviders"`
  canEdit                       bool         `json "canEdit"`
  canShare                      bool         `json "canShare"`
  canViewPrivateItems           bool         `json "canViewPrivateItems"`
  changeKey                     string       `json "changeKey"`
  color                         string       `json "color"`
  defaultOnlineMeetingProvider  string       `json "defaultOnlineMeetingProvider"`
  id                            string       `json "id"`
  isRemovable                   bool         `json "isRemovable"`
  isTallyingResponses           bool         `json "isTallyingResponses"`
  name                          string       `json "name"`
  owner                         EmailAddress `json "owner"`
}
