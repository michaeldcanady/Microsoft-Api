package outlook

import(
  "strings"
  "fmt"
)

//Filter is a item to base a filter on
func Filter(str string) string {
	return "$filter="+strings.ToLower(str)
}

func Eq(str string) string {
  return fmt.Sprintf("eq '%s'",str)
}

func Ne(str string) string {
  return fmt.Sprintf("ne '%s'",str)
}

func Lt(str string) string {
  return fmt.Sprintf("lt '%s'",str)
}

func Gt(str string) string {
  return fmt.Sprintf("gt '%s'",str)
}

func Le(str string) string {
  return fmt.Sprintf("le '%s'",str)
}

func Ge(str string) string {
  return fmt.Sprintf("ge '%s'",str)
}

func AND() string{
  return "and"
}

func URLAnd() string{
  return "&"
}

func OR() string{
  return "or"
}

// NOT SURE HOW TO USE
func In(){

}

// NOT SURE HOW TO USE
func Not(){

}

// NOT SURE HOW TO USE
func Any(){

}

// NOT SURE HOW TO USE
func All(){

}

func StartsWith(parameter, value string) string{
  return fmt.Sprintf("statsWith(%s,'%s')",parameter,value)
}

func EndsWith(parameter, value string) string{
  return fmt.Sprintf("endsWith(%s,'%s')",parameter,value)
}

func OrderBy(sort string,value ...string) string{
  return fmt.Sprintf("$orderby=%s%s%s",strings.Join(value,"/"),"%20%",sort)
}

func AndOrder(sort string,value ...string) string{
  return fmt.Sprintf(",%s%s%s",strings.Join(value,"/"),"%20%",sort)
}

func Search(value string,location ...string) string{
  if len(location) < 1{
    return fmt.Sprintf("$search='%s'",value)
  }else{
    return fmt.Sprintf("$search='%s:%s'",strings.Join(location,""),value)
  }
}
