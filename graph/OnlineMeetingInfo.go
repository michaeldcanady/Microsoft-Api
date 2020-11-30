package graph

import(

)

type OnlineMeetingInfo struct{
  ConferenceId    string   `json "conferenceId"`
  JoinUrl         string   `json "joinUrl"`
  Phones          []phone  `json "phones"`
  QuickDial       string   `json "quickDial"`
  TollFreeNumbers []string `json "tollFreeNumbers"`
  TollNumber      string   `json "tollNumber"`
}
