package graph

import(

)

type WorkHours struct{
  Day       string         `json "day"`
  TimeSlots []WorkTimeSlot `json "timeSlotss"`
}
