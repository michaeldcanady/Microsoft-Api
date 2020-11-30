package graph

import(
  "time"
)

type Contact struct{
  AssistantName        string          `json"assistantName"`
  Birthday             time.Time       `json"birthday"`
  BusinessAddress      PhysicalAddress `json"businessAddress"`
  BusinessHomePage     string          `json"businessHomePage"`
  BusinessPhones       []string        `json"businessPhones"`
  Categories           []string        `json"categories"`
  ChangeKey            string          `json"changeKey"`
  Children             []string        `json"children"`
  CompanyName          string          `json"companyName"`
  CreatedDateTime      time.Time       `json"createdDateTime"`
  Department           string          `json"department"`
  DisplayName          string          `json"displayName"`
  EmailAddresses       EmailAddress    `json"emailAddresses"`
  FileAs               string          `json"fileAs"`
  Generation           string          `json"generation"`
  GivenName            string          `json"givenName"`
  HomeAddress          PhysicalAddress `json"homeAddress"`
  HomePhones           []string        `json"homePhones"`
  Id                   string          `json"id"`
  ImAddresses          []string        `json"imAddresses"`
  Initials             string          `json"initials"`
  JobTitle             string          `json"jobTitle"`
  LastModifiedDateTime time.Time       `json"lastModifiedDateTime"`
  Manager              string          `json"manager"`
  MiddleName           string          `json"middleName"`
  MobilePhone          string          `json"mobilePhone"`
  NickName             string          `json"nickName"`
  OfficeLocation       string          `json"officeLocation"`
  OtherAddress         PhysicalAddress `json"otherAddress"`
  ParentFolderId       string          `json"parentFolderId"`
  PersonalNotes        string          `json"personalNotes"`
  Profession           string          `json"profession"`
  SpouseName           string          `json"spouseName"`
  Surname              string          `json"surname"`
  Title                string          `json"title"`
  YomiCompanyName      string          `json"yomiCompanyName"`
  YomiGivenName        string          `json"yomiGivenName"`
  YomiSurname          string          `json"yomiSurname"`
  Photo                ProfilePhoto    `json"photo"`
}
