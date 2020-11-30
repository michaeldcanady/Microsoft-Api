package graph

import(

)

type PhysicalAddress struct{
  City            string `json"city"`
  CountryOrRegion string `json"countryOrRegion"`
  PostalCode      string `json"postalCode"`
  State           string `json"state"`
  Street          string `json"street"`
}
