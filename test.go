package main

import(
  //"github.com/michaeldcanady/GoOutlook/GoOutlook"
  "fmt"
)


func main(){
  client := NewClient("dmcanady","**************")
  fmt.Println(client.PULL("v1.0","me/drive",Select("from","subject")))
}
