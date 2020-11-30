package booking

import(

)

type Business struct{
  address {"@odata.type": "microsoft.graph.physicalAddress"},
  businessHours [{"@odata.type": "microsoft.graph.bookingWorkHours"}],
  businessType string
  defaultCurrencyIso string
  displayName string
  email string
  id string
  isPublished bool
  phone String
  publicUrl string
  schedulingPolicy {"@odata.type": "microsoft.graph.bookingSchedulingPolicy"},
  webSiteUrl string
}
