package main

import(
  //"github.com/michaeldcanady/GoOutlook/GoOutlook"
  "fmt"
  //"strings"
)

var(
  OUTLOOK_APP_ID = "98a16ce9-8f10-46df-a8c7-0c8b19d73e39"
  OUTLOOK_APP_SECRET = "BwhkGra8B.32.j48B~i2uAE~PwuyY0L_Ac"
)

func main(){

  app := SetClientAppID(OUTLOOK_APP_ID)
  secret := SetClientAppSecret(OUTLOOK_APP_SECRET)
  redirect := SetClientRedirectURI("https://snappy.appypie.com/app/onedrive")

  client,err := NewClient(app,secret,redirect)
  if err != nil{
    panic(err)
  }

  client.GetAccessToken()
  URL := ""
  fmt.Println("Please past the url you were redirected to here: ")
  fmt.Scanln(&URL)
  session,err := NewSession(client,URL)
  if err != nil {
    panic(err)
  }
  fmt.Println(session)

//  session,err := NewSession(client,"")
//  if err != nil {
//    panic(err)
//  }
//  fmt.Println(session)

  //fmt.Println(client.client)
	//fmt.Println(client.baseURL)
	//fmt.Println(client.userAgent)
  //fmt.Println(client.mediaType)
	//fmt.Println(client.appID)
	//fmt.Println(client.appSecret)
	//fmt.Println(client.redirectURI)
	//fmt.Println(client.scope)
}
