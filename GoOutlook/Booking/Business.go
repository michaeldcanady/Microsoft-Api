package main

import(

)

type Business struct{
  address PhysicalAddress
  businessHours []WorkHours
  businessType string
  defaultCurrencyIso string
  displayName string
  email string
  id string
  isPublished bool
  phone string
  publicUrl string
  schedulingPolicy SchedulingPolicy
  webSiteUrl string
}

func (b *Business) LIST() []Business{

}

func (b *Business) CREATE() Business{

}

func (b *Business) GET() Business{

}

func (b *Business) UPDATE() Business{

}

func (b *Business) DELETE(){

}
