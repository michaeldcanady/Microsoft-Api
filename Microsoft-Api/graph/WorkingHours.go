package graph

imort(

)

type WorkingHours struct{
  DaysOfWeek []string     `json "daysOfWeek"`
  StartTime  time.Time    `json "startTime"`
  EndTime    time.Time    `json "endTime"`
  TimeZone   TimeZoneBase `json "timeZone"`
}
