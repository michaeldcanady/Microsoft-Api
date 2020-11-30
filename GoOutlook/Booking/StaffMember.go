package main

import(

)

type StaffMember struct{
  "availabilityIsAffectedByPersonalCalendar": true,
  "colorIndex": 1024,
  "displayName": "String",
  "emailAddress": "String",
  "id": "String (identifier)",
  "role": "string",
  "useBusinessHours": true,
  "workingHours": [{"@odata.type": "microsoft.graph.bookingWorkHours"}]
}

func (s *StaffMember) LIST() []StaffMember{

}

func (s *StaffMember) CREATE() StaffMember{

}

func (s *StaffMember) GET() StaffMember{

}

func (s *StaffMember) UPDATE() StaffMember{

}

func (s *StaffMember) DELETE() StaffMember{

}
