package booking

import(
  "time"
)

type Appointment struct{
  customerEmailAddress string `json: "customerEmailAddress"`
  customerId string `json: "customerId"`
  customerLocation time.Location `json: "customerLocation"`
  customerName string `json: "customerName"`
  customerNotes string `json: "customerNotes"`
  customerPhone string `json: "customerPhone"`
  duration time.Duration `json: "duration"`
  end time.Time `json: "end"`
  id string `json: "id"`
  invoiceAmount int `json: "invoiceAmount"`
  invoiceDate time.Time `json: "invoiceDate"`
  invoiceId string `json: "invoiceId"`
  invoiceStatus string `json: "invoiceStatus"`
  invoiceUrl string `json: "invoiceUrl"`
  optOutOfCustomerEmail bool `json: "optOutOfCustomerEmail"`
  postBuffer time.Duration `json: "postBuffer"`
  preBuffer time.Duration `json: "preBuffer"`
  price int `json: "price"`
  priceType string `json: "priceType"`
  reminders []string `json: "reminders"`
  selfServiceAppointmentId string `json: "selfServiceAppointmentId"`
  serviceId string `json: "serviceId"`
  serviceLocation time.Location `json: "serviceLocation"`
  serviceName string `json: "serviceName"`
  serviceNotes string `json: "serviceNotes"`
  staffMemberIds []string `json: "staffMemberIds"`
  start time.Time `json: "start"`
}

// Gets a list of all booking Appointments
func (a *Appointment) LIST()(){

}

func (a *Appointment) CREATE()(){

}

func (a *Appointment) GET()(){

}

func (a *Appointment) UPDATE()(){

}

func (a *Appointment) DELETE()(){

}

func (a *Appointment) CANCEL()(){

}
