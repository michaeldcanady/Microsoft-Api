package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)


var(
  URL = "https://graph.microsoft.com/"
)

type Client struct{
  Username string
  Password string
}

func NewClient(username,password string) Client{
  return Client{Username:username,Password:password}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}
var httpClient = &http.Client{
	Transport: tr,
	Timeout:   time.Duration(10 * time.Second),
}

type errStruct struct {
	Err    string `json:"error"`
	Reason string
}

func (e errStruct) Error() string {
	if e.Reason == "" {
		return e.Err
	}
	return fmt.Sprintf("%e: %s", e.Err, e.Reason)
}

func (C *Client) PULL(version,resource,opts string)([]map[string]string, int){
  bufDT := &bytes.Buffer{}

  address := fmt.Sprintf("%s%s/%s?%s",URL,version,resource,opts)

  reqDT, err := http.NewRequest(http.MethodGet, address, bufDT)
	checkErr(err)

  reqDT.SetBasicAuth(C.Username, C.Password)

  resDT, err := httpClient.Do(reqDT)
	checkErr(err)

  bufDT.Reset()

  var echeck errStruct

  err = json.NewDecoder(io.TeeReader(resDT.Body, bufDT)).Decode(&echeck)
  fmt.Println(err)
	checkErr(err)

  var dT, dF struct {
		Records []map[string]string
	}

  json.NewDecoder(bufDT).Decode(&dT)

  for i, eachDT := range dT.Records {
		for k, v := range eachDT {
			if strings.Contains(k, "-ID") {
				continue
			}
			if dF.Records[i][k] != v {
				dT.Records[i][k+"-ID"] = dF.Records[i][k]
			}
		}
	}

	return dT.Records, len(dT.Records)
}
