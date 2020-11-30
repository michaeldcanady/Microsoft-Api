package outlook

import(

)

type booking struct{
  customerEmailAddress String
  customerId String
  customerLocation date.Location
  customerName String
  customerNotes String
  customerPhone String
  duration String (timestamp)
  end: {"@odata.type": "microsoft.graph.dateTimeTimeZone"}
  id: String (identifier)
  invoiceAmount int
  invoiceDate {"@odata.type": "microsoft.graph.dateTimeTimeZone"}
  invoiceId String
  invoiceStatus string
  invoiceUrl string
  optOutOfCustomerEmail bool
  postBuffer String (timestamp)
  preBuffer String (timestamp)
  price int
  priceType string
  reminders [{"@odata.type": "microsoft.graph.bookingReminder"}]
  selfServiceAppointmentId string
  serviceId string
  serviceLocation {"@odata.type": "microsoft.graph.location"}
  serviceName string
  serviceNotes string
  staffMemberIds ["String"]
  start {"@odata.type": "microsoft.graph.dateTimeTimeZone"}
}
