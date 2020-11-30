package outlook

import(

)


var(
  BETAURL = "https://graph.microsoft.com/beta/{resource}?[query_parameters]"
  VERSION = "v1.0"
  URL = "https://graph.microsoft.com/"+VERSION+"/{resource}?[query_parameters]"
)

type Client struct{
  Username string
  Password string
}

func NewClient(username,password string) Client{
  return Client{Username:username,Password:password}
}

func (C *Client) PULL() string{
  return URL
}
