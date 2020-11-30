package graph

import(

)

type AssignedLicense struct{
  DisabledPlans []string `json "disabledPlans"`
  SkuId         string   `json "skuId"`
}
