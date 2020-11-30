package main

import(
  "github.com/michaeldcanady/GoOutlook/GoOutlook"
  "fmt"
)


func main(){
  client := outlook.NewClient("dmcanady","********")
  fmt.Println(client.PULL())
  fmt.Println(outlook.Search("Jimmy Fallon",))
}
