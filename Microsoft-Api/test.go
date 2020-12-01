package main

import(
  //"github.com/michaeldcanady/GoOutlook/GoOutlook"
  "fmt"
)

var(
  OUTLOOK_APP_ID = "98a16ce9-8f10-46df-a8c7-0c8b19d73e39"
  OUTLOOK_APP_SECRET = "80d93fe8-71df-49c3-87ae-f8a9bc6a2ad8"
)

func main(){

  app := SetClientAppID(OUTLOOK_APP_ID)
  secret := SetClientAppSecret(OUTLOOK_APP_SECRET)

  client,err := NewClient(app,secret)
  if err != nil{
    panic(err)
  }

//  session,err := NewSession(client,"")
//  if err != nil {
//    panic(err)
//  }
//  fmt.Println(session)

  fmt.Println(client.client)
	fmt.Println(client.baseURL)
	fmt.Println(client.userAgent)
  fmt.Println(client.mediaType)
	fmt.Println(client.appID)
	fmt.Println(client.appSecret)
	fmt.Println(client.redirectURI)
	fmt.Println(client.scope)
}
