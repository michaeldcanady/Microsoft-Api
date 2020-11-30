package users

import(
  "github.com/michaeldcanady/graph"
  "time"
)

type User struct{
  aboutMe string
  accountEnabled bool `json "accountEnabled"`
  ageGroup string
  assignedLicenses graph.AssignedLicense
  assignedPlans graph.AssignedPlan
  birthday time.Time
  businessPhones []string
  city string
  companyName string
  consentProvidedForMinor string
  country string
  createdDateTime string `json "createdDateTime"`
  creationType string
  deletedDateTime time.Time
  department string
  displayName string
  employeeHireDate string `json "employeeHireDate"`
  employeeId string
  employeeOrgData graph.EmployeeOrgData
  employeeType string
  externalUserState string `json "externalUserState"`
  externalUserStateChangeDateTime string `json "externalUserStateChangeDateTime"`
  faxNumber string
  givenName string
  hireDate time.Time
  id string
  identities graph.ObjectIdentity
  interests []string
  isResourceAccount bool
  jobTitle string
  legalAgeGroupClassification string
  licenseAssignmentStates graph.LicenseAssignmentState
  mail string
  mailboxSettings": {"@odata.type": "microsoft.graph.mailboxSettings"},
  mailNickname string
  mobilePhone string
  mySite string
  officeLocation string
  onPremisesDistinguishedName string
  onPremisesDomainName string
  onPremisesExtensionAttributes": {"@odata.type": "microsoft.graph.onPremisesExtensionAttributes"},
  onPremisesImmutableId string
  onPremisesLastSyncDateTime time.Time
  onPremisesProvisioningErrors": [{"@odata.type": "microsoft.graph.onPremisesProvisioningError"}],
  onPremisesSamAccountName string
  onPremisesSecurityIdentifier string
  onPremisesSyncEnabled bool
  onPremisesUserPrincipalName string
  otherMails []string
  passwordPolicies string
  passwordProfile": {"@odata.type": "microsoft.graph.passwordProfile"},
  pastProjects []string
  postalCode string
  preferredDataLocation string
  preferredLanguage string
  preferredName string
  provisionedPlans": [{"@odata.type": "microsoft.graph.provisionedPlan"}],
  proxyAddresses []string
  refreshTokensValidFromDateTime": "2019-02-07T21:53:13.084Z",
  responsibilities []string
  schools []string
  showInAddressList bool
  signInSessionsValidFromDateTime": "2019-02-07T21:53:13.084Z",
  skills []string
  state string
  streetAddress string
  surname string
  usageLocation string
  userPrincipalName string
  userType string
  calendar": {"@odata.type": "microsoft.graph.calendar"},
  calendarGroups": [{"@odata.type": "microsoft.graph.calendarGroup"}],
  calendarView": [{"@odata.type": "microsoft.graph.event"}],
  calendars": [{"@odata.type": "microsoft.graph.calendar"}],
  contacts": [{"@odata.type": "microsoft.graph.contact"}],
  contactFolders": [{"@odata.type": "microsoft.graph.contactFolder"}],
  createdObjects": [{"@odata.type": "microsoft.graph.directoryObject"}],
  directReports": [{"@odata.type": "microsoft.graph.directoryObject"}],
  drive": {"@odata.type": "microsoft.graph.drive"},
  drives": [{"@odata.type": "microsoft.graph.drive"}],
  insights": {"@odata.type": "microsoft.graph.itemInsights"},
  settings": {"@odata.type": "microsoft.graph.userSettings"},
  events": [{"@odata.type": "microsoft.graph.event"}],
  extensions": [{"@odata.type": "microsoft.graph.extension"}],
  inferenceClassification": {"@odata.type": "microsoft.graph.inferenceClassification"},
  mailFolders": [{"@odata.type": "microsoft.graph.mailFolder"}],
  manager": {"@odata.type": "microsoft.graph.directoryObject"},
  memberOf": [{"@odata.type": "microsoft.graph.directoryObject"}],
  joinedTeams": [{"@odata.type": "microsoft.graph.group"}],
  teamwork": {"@odata.type": "microsoft.graph.teamwork"},
  messages": [{ "@odata.type": "microsoft.graph.message"}],
  outlook": {"@odata.type": "microsoft.graph.outlookUser"},
  ownedDevices": [{"@odata.type": "microsoft.graph.directoryObject"}],
  photo": {"@odata.type": "microsoft.graph.profilePhoto"},
  profile": {"@odata.type": "microsoft.graph.profile"},
  registeredDevices": [{"@odata.type": "microsoft.graph.directoryObject"}],
  signInActivity": {"@odata.type": "microsoft.graph.signInActivity"}
}
