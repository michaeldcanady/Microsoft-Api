package main

import(
  "time"
  "github.com/michaeldcanady/GoOutlook/GoOutlook"
)

type Appointment struct{
  customerEmailAddress string `json: "customerEmailAddress"`
  customerId string `json: "customerId"`
  customerLocation Location `json: "customerLocation"`
  customerName string `json: "customerName"`
  customerNotes string `json: "customerNotes"`
  customerPhone string `json: "customerPhone"`
  duration time.Duration `json: "duration"`
  end outlook.TimeZone `json: "end"`
  id string `json: "id"`
  invoiceAmount int `json: "invoiceAmount"`
  invoiceDate outlook.TimeZone `json: "invoiceDate"`
  invoiceId string `json: "invoiceId"`
  invoiceStatus string `json: "invoiceStatus"`
  invoiceUrl string `json: "invoiceUrl"`
  optOutOfCustomerEmail bool `json: "optOutOfCustomerEmail"`
  postBuffer time.Duration `json: "postBuffer"`
  preBuffer time.Duration `json: "preBuffer"`
  price int `json: "price"`
  priceType string `json: "priceType"`
  reminders []Reminder `json: "reminders"`
  selfServiceAppointmentId string `json: "selfServiceAppointmentId"`
  serviceId string `json: "serviceId"`
  serviceLocation Location `json: "serviceLocation"`
  serviceName string `json: "serviceName"`
  serviceNotes string `json: "serviceNotes"`
  staffMemberIds []string `json: "staffMemberIds"`
  start outlook.TimeZone `json: "start"`
}

// Gets a list of all booking Appointments
func (a *Appointment) LIST() []Appointment{

}

func (a *Appointment) CREATE() Appointment{

}

func (a *Appointment) GET() Appointment{

}

func (a *Appointment) UPDATE() Appointment{

}

func (a *Appointment) DELETE(){

}

func (a *Appointment) CANCEL(){

}
