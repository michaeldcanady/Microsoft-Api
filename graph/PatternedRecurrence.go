package graph

import(

)

type PatternedRecurrence struct{
  Pattern RecurrencePattern `json "pattern"`
  Range   RecurrencePattern `json "range"`
}
