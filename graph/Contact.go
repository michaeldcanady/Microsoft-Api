package graph

import(
  "time"
)

type Contact struct{
  assistantName        string          `json"assistantName"`
  birthday             time.Time       `json"birthday"`
  businessAddress      PhysicalAddress `json"businessAddress"`
  businessHomePage     string          `json"businessHomePage"`
  businessPhones       []string        `json"businessPhones"`
  categories           []string        `json"categories"`
  changeKey            string          `json"changeKey"`
  children             []string        `json"children"`
  companyName          string          `json"companyName"`
  createdDateTime      time.Time       `json"createdDateTime"`
  department           string          `json"department"`
  displayName          string          `json"displayName"`
  emailAddresses       EmailAddress    `json"emailAddresses"`
  fileAs               string          `json"fileAs"`
  generation           string          `json"generation"`
  givenName            string          `json"givenName"`
  homeAddress          PhysicalAddress `json"homeAddress"`
  homePhones           []string        `json"homePhones"`
  id                   string          `json"id"`
  imAddresses          []string        `json"imAddresses"`
  initials             string          `json"initials"`
  jobTitle             string          `json"jobTitle"`
  lastModifiedDateTime time.Time       `json"lastModifiedDateTime"`
  manager              string          `json"manager"`
  middleName           string          `json"middleName"`
  mobilePhone          string          `json"mobilePhone"`
  nickName             string          `json"nickName"`
  officeLocation       string          `json"officeLocation"`
  otherAddress         PhysicalAddress `json"otherAddress"`
  parentFolderId       string          `json"parentFolderId"`
  personalNotes        string          `json"personalNotes"`
  profession           string          `json"profession"`
  spouseName           string          `json"spouseName"`
  surname              string          `json"surname"`
  title                string          `json"title"`
  yomiCompanyName      string          `json"yomiCompanyName"`
  yomiGivenName        string          `json"yomiGivenName"`
  yomiSurname          string          `json"yomiSurname"`
  photo                ProfilePhoto    `json"photo"`
}
