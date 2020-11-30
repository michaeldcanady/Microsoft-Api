package main

import(

)

type StaffMember struct{
  AvailabilityIsAffectedByPersonalCalendar bool
  colorIndex int
  displayName string
  emailAddress string
  id string
  role string
  useBusinessHours bool
  workingHours WorkHours
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
