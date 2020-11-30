package main

import(
  "time"
)

type Service struct {
  defaultDuration time.Time
  defaultLocation Location
  defaultPrice int
  defaultPriceType string
  defaultReminders Reminder
  description string
  displayName string
  emailAddress string
  id string
  isHiddenFromCustomers bool
  notes string
  postBuffer time.Time
  preBuffer time.Time
  schedulingPolicy SchedulingPolicy
  staffMemberIds []string
}

func (s *Service) LIST() []Service{

}

func (s *Service) CREATE() Service{

}

func (s *Service) GET() Service{

}

func (s *Service) UPDATE() Service{

}

func (s *Service) DELETE() Service{

}
