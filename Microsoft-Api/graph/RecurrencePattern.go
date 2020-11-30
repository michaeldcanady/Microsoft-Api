package graph

import(

)

type RecurrencePattern struct{
  DayOfMonth     int      `json "dayOfMonth"`
  DaysOfWeek     []string `json "daysOfWeek"`
  FirstDayOfWeek string   `json "firstDayOfWeek"`
  Index          string   `json "index"`
  Interval       int      `json "interval"`
  Month          int      `json "month"`
  Type           string   `json "type"`
}
