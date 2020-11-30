package main

import(

)

type Customer struct{
  DisplayName string
  EmailAddress string
  Id string
}

func (c *Customer) LIST() []Customer{
  URL = "https://graph.microsoft.com/beta/bookingBusinesses"
}

func (c *Customer) CREATE() Customer{

}

func (c *Customer) GET() Customer{

}

func (c *Customer) UPDATE() Customer{

}

func (c *Customer) DELETE() {

}
