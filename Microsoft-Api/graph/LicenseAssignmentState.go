package graph

import(

)

type LicenseAssignmentState struct{
  AssignedByGroup string `json "assignedByGroup"`
  DisabledPlans   string `json "disabledPlans"`
  Error           string `json "error"`
  SkuId           string `json "skuId"`
  State           string `json "state"`
}
